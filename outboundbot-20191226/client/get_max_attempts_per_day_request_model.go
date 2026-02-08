// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaxAttemptsPerDayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *GetMaxAttemptsPerDayRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *GetMaxAttemptsPerDayRequest
	GetStrategyLevel() *int32
}

type GetMaxAttemptsPerDayRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 2bfa5ae4-7185-4227-a3b8-328f26f11be1
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// Policy level. The default value is 2 for business instances.
	//
	// - 0: System
	//
	// - 1: Tenant
	//
	// - 2: Instance
	//
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s GetMaxAttemptsPerDayRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMaxAttemptsPerDayRequest) GoString() string {
	return s.String()
}

func (s *GetMaxAttemptsPerDayRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *GetMaxAttemptsPerDayRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *GetMaxAttemptsPerDayRequest) SetEntryId(v string) *GetMaxAttemptsPerDayRequest {
	s.EntryId = &v
	return s
}

func (s *GetMaxAttemptsPerDayRequest) SetStrategyLevel(v int32) *GetMaxAttemptsPerDayRequest {
	s.StrategyLevel = &v
	return s
}

func (s *GetMaxAttemptsPerDayRequest) Validate() error {
	return dara.Validate(s)
}
