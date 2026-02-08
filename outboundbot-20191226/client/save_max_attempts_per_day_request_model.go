// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveMaxAttemptsPerDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *SaveMaxAttemptsPerDayRequest
	GetEntryId() *string
	SetMaxAttemptsPerDay(v int32) *SaveMaxAttemptsPerDayRequest
	GetMaxAttemptsPerDay() *int32
	SetStrategyLevel(v int32) *SaveMaxAttemptsPerDayRequest
	GetStrategyLevel() *int32
}

type SaveMaxAttemptsPerDayRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c8bf820a-6a8a-47bc-99bf-97593df8faa8
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// Maximum number of redial attempts per day. Default value is 3 if not specified.
	//
	// example:
	//
	// 3
	MaxAttemptsPerDay *int32 `json:"MaxAttemptsPerDay,omitempty" xml:"MaxAttemptsPerDay,omitempty"`
	// Policy level (Required)
	//
	// - 2: instance
	//
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s SaveMaxAttemptsPerDayRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveMaxAttemptsPerDayRequest) GoString() string {
	return s.String()
}

func (s *SaveMaxAttemptsPerDayRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *SaveMaxAttemptsPerDayRequest) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *SaveMaxAttemptsPerDayRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *SaveMaxAttemptsPerDayRequest) SetEntryId(v string) *SaveMaxAttemptsPerDayRequest {
	s.EntryId = &v
	return s
}

func (s *SaveMaxAttemptsPerDayRequest) SetMaxAttemptsPerDay(v int32) *SaveMaxAttemptsPerDayRequest {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *SaveMaxAttemptsPerDayRequest) SetStrategyLevel(v int32) *SaveMaxAttemptsPerDayRequest {
	s.StrategyLevel = &v
	return s
}

func (s *SaveMaxAttemptsPerDayRequest) Validate() error {
	return dara.Validate(s)
}
