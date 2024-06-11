package twikey

import "time"

type SubscriptionState string

const (
	SubscriptionStateActive    SubscriptionState = "active"
	SubscriptionStateSuspended SubscriptionState = "suspended"
	SubscriptionStateCancelled SubscriptionState = "cancelled"
	SubscriptionStateClosed    SubscriptionState = "closed"
)

type SubscriptionRecurrence string

const (
	SubscriptionRecurrenceWeekly    SubscriptionRecurrence = "1w"
	SubscriptionRecurrenceMonthly   SubscriptionRecurrence = "1m"
	SubscriptionRecurrenceBiMonthly SubscriptionRecurrence = "2m"
	SubscriptionRecurrenceQuarterly SubscriptionRecurrence = "3m"
	SubscriptionRecurrenceTrimester SubscriptionRecurrence = "4m"
	SubscriptionRecurrenceSemester  SubscriptionRecurrence = "6m"
	SubscriptionRecurrenceAnnual    SubscriptionRecurrence = "12m"
)

type Subscription struct {
	Id      int64             `json:"id"`
	State   SubscriptionState `json:"state"`
	Amount  float64           `json:"amount"`
	Message string            `json:"message"`
	Ref     string            `json:"ref"`
	Plan    int64             `json:"plan,omitempty"`
	Runs    int               `json:"runs"`
	// StopAfter contains the maximum amount of runs a subscription can have before it gets closed.
	StopAfter        int                    `json:"stopAfter,omitempty"`
	Start            time.Time              `json:"start"`
	Last             time.Time              `json:"last"`
	Next             time.Time              `json:"next"`
	Recurrence       SubscriptionRecurrence `json:"recurrence"`
	MandateReference string                 `json:"mndtId"`
}
