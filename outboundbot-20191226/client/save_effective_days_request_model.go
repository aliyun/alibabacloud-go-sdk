// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEffectiveDaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveDays(v int32) *SaveEffectiveDaysRequest
	GetEffectiveDays() *int32
	SetEntryId(v string) *SaveEffectiveDaysRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *SaveEffectiveDaysRequest
	GetStrategyLevel() *int32
}

type SaveEffectiveDaysRequest struct {
	// Effective period
	//
	// example:
	//
	// 30
	EffectiveDays *int32 `json:"EffectiveDays,omitempty" xml:"EffectiveDays,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e90b5b8e-c8b4-4182-b28d-a5aa81685e49
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// Policy level
	//
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s SaveEffectiveDaysRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveEffectiveDaysRequest) GoString() string {
	return s.String()
}

func (s *SaveEffectiveDaysRequest) GetEffectiveDays() *int32 {
	return s.EffectiveDays
}

func (s *SaveEffectiveDaysRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *SaveEffectiveDaysRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *SaveEffectiveDaysRequest) SetEffectiveDays(v int32) *SaveEffectiveDaysRequest {
	s.EffectiveDays = &v
	return s
}

func (s *SaveEffectiveDaysRequest) SetEntryId(v string) *SaveEffectiveDaysRequest {
	s.EntryId = &v
	return s
}

func (s *SaveEffectiveDaysRequest) SetStrategyLevel(v int32) *SaveEffectiveDaysRequest {
	s.StrategyLevel = &v
	return s
}

func (s *SaveEffectiveDaysRequest) Validate() error {
	return dara.Validate(s)
}
