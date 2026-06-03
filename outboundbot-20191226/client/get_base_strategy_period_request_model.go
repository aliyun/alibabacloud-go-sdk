// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaseStrategyPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *GetBaseStrategyPeriodRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *GetBaseStrategyPeriodRequest
	GetStrategyLevel() *int32
}

type GetBaseStrategyPeriodRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 024f8cf0-c842-4c01-b74b-c8667e4579c7
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// example:
	//
	// 2
	StrategyLevel *int32 `json:"StrategyLevel,omitempty" xml:"StrategyLevel,omitempty"`
}

func (s GetBaseStrategyPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBaseStrategyPeriodRequest) GoString() string {
	return s.String()
}

func (s *GetBaseStrategyPeriodRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *GetBaseStrategyPeriodRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *GetBaseStrategyPeriodRequest) SetEntryId(v string) *GetBaseStrategyPeriodRequest {
	s.EntryId = &v
	return s
}

func (s *GetBaseStrategyPeriodRequest) SetStrategyLevel(v int32) *GetBaseStrategyPeriodRequest {
	s.StrategyLevel = &v
	return s
}

func (s *GetBaseStrategyPeriodRequest) Validate() error {
	return dara.Validate(s)
}
