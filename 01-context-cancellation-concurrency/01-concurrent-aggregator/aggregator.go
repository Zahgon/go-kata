package concurrent_aggregator

import (
	"context"
	"log/slog"
	"time"

	"github.com/medunes/go-kata/01-context-cancellation-concurrency/01-concurrent-aggregator/order"
	"github.com/medunes/go-kata/01-context-cancellation-concurrency/01-concurrent-aggregator/profile"
)

type Option func(ua *UserAggregator)

func WithLogger(l *slog.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

type UserAggregator struct {
	orderService   order.Service
	profileService profile.Service
	timeout        time.Duration
	logger         *slog.Logger
}

func NewUserAggregator(os order.Service, ps profile.Service, opts ...Option) *UserAggregator {
	_ = "STUB: not implemented"
	return nil
}

// avoid nil panics, default writes to stderr

type AggregatedProfile struct {
	Name string
	Cost float64
}

func (ua *UserAggregator) Aggregate(ctx context.Context, id int) ([]*AggregatedProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
