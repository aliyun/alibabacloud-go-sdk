// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyStrategyIdFilter interface {
	dara.Model
	String() string
	GoString() string
	SetEq(v string) *NotifyStrategyIdFilter
	GetEq() *string
}

type NotifyStrategyIdFilter struct {
	// The exact ID of the notification strategy to retrieve.
	Eq *string `json:"eq,omitempty" xml:"eq,omitempty"`
}

func (s NotifyStrategyIdFilter) String() string {
	return dara.Prettify(s)
}

func (s NotifyStrategyIdFilter) GoString() string {
	return s.String()
}

func (s *NotifyStrategyIdFilter) GetEq() *string {
	return s.Eq
}

func (s *NotifyStrategyIdFilter) SetEq(v string) *NotifyStrategyIdFilter {
	s.Eq = &v
	return s
}

func (s *NotifyStrategyIdFilter) Validate() error {
	return dara.Validate(s)
}
