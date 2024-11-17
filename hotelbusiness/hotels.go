//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	exchangeGuests := make(map[int]int)

	for _, guest := range guests {
		exchangeGuests[guest.CheckInDate] += 1
		exchangeGuests[guest.CheckOutDate] -= 1
	}

	dates := make([]int, 0, len(exchangeGuests))
	for date := range exchangeGuests {
		dates = append(dates, date)
	}
	sort.Ints(dates)

	loads := make([]Load, 0)
	currentGuestCount := 0

	for _, date := range dates {
		guestCountChange := exchangeGuests[date]
		currentGuestCount += guestCountChange
		if len(loads) > 0 && loads[len(loads)-1].GuestCount == currentGuestCount {
			continue
		}
		loads = append(loads, Load{date, currentGuestCount})
	}

	return loads
}
