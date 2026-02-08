// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmptyNumberNoMoreCallsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntryId(v string) *GetEmptyNumberNoMoreCallsInfoRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *GetEmptyNumberNoMoreCallsInfoRequest
	GetStrategyLevel() *int32
}

type GetEmptyNumberNoMoreCallsInfoRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// db3e679b-7d5e-4d9b-828a-345adca455f3
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

func (s GetEmptyNumberNoMoreCallsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmptyNumberNoMoreCallsInfoRequest) GoString() string {
	return s.String()
}

func (s *GetEmptyNumberNoMoreCallsInfoRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *GetEmptyNumberNoMoreCallsInfoRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *GetEmptyNumberNoMoreCallsInfoRequest) SetEntryId(v string) *GetEmptyNumberNoMoreCallsInfoRequest {
	s.EntryId = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoRequest) SetStrategyLevel(v int32) *GetEmptyNumberNoMoreCallsInfoRequest {
	s.StrategyLevel = &v
	return s
}

func (s *GetEmptyNumberNoMoreCallsInfoRequest) Validate() error {
	return dara.Validate(s)
}
