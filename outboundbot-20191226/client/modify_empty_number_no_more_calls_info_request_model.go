// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmptyNumberNoMoreCallsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmptyNumberNoMoreCalls(v bool) *ModifyEmptyNumberNoMoreCallsInfoRequest
	GetEmptyNumberNoMoreCalls() *bool
	SetEntryId(v string) *ModifyEmptyNumberNoMoreCallsInfoRequest
	GetEntryId() *string
	SetStrategyLevel(v int32) *ModifyEmptyNumberNoMoreCallsInfoRequest
	GetStrategyLevel() *int32
}

type ModifyEmptyNumberNoMoreCallsInfoRequest struct {
	// Indicates whether to stop outbound calls to nonexistent numbers
	//
	// example:
	//
	// true
	EmptyNumberNoMoreCalls *bool `json:"EmptyNumberNoMoreCalls,omitempty" xml:"EmptyNumberNoMoreCalls,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
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

func (s ModifyEmptyNumberNoMoreCallsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmptyNumberNoMoreCallsInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) GetEmptyNumberNoMoreCalls() *bool {
	return s.EmptyNumberNoMoreCalls
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) GetStrategyLevel() *int32 {
	return s.StrategyLevel
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) SetEmptyNumberNoMoreCalls(v bool) *ModifyEmptyNumberNoMoreCallsInfoRequest {
	s.EmptyNumberNoMoreCalls = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) SetEntryId(v string) *ModifyEmptyNumberNoMoreCallsInfoRequest {
	s.EntryId = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) SetStrategyLevel(v int32) *ModifyEmptyNumberNoMoreCallsInfoRequest {
	s.StrategyLevel = &v
	return s
}

func (s *ModifyEmptyNumberNoMoreCallsInfoRequest) Validate() error {
	return dara.Validate(s)
}
