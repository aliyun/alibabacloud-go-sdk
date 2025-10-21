// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogReservePolicy interface {
	dara.Model
	String() string
	GoString() string
	SetExpirationDays(v int64) *LogReservePolicy
	GetExpirationDays() *int64
	SetOpenHistory(v bool) *LogReservePolicy
	GetOpenHistory() *bool
}

type LogReservePolicy struct {
	// example:
	//
	// 7
	ExpirationDays *int64 `json:"expirationDays,omitempty" xml:"expirationDays,omitempty"`
	// example:
	//
	// true
	OpenHistory *bool `json:"openHistory,omitempty" xml:"openHistory,omitempty"`
}

func (s LogReservePolicy) String() string {
	return dara.Prettify(s)
}

func (s LogReservePolicy) GoString() string {
	return s.String()
}

func (s *LogReservePolicy) GetExpirationDays() *int64 {
	return s.ExpirationDays
}

func (s *LogReservePolicy) GetOpenHistory() *bool {
	return s.OpenHistory
}

func (s *LogReservePolicy) SetExpirationDays(v int64) *LogReservePolicy {
	s.ExpirationDays = &v
	return s
}

func (s *LogReservePolicy) SetOpenHistory(v bool) *LogReservePolicy {
	s.OpenHistory = &v
	return s
}

func (s *LogReservePolicy) Validate() error {
	return dara.Validate(s)
}
