// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateWebLockFileEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDealAll(v int32) *OperateWebLockFileEventsRequest
	GetDealAll() *int32
	SetEventIds(v []*int64) *OperateWebLockFileEventsRequest
	GetEventIds() []*int64
	SetOperationCode(v string) *OperateWebLockFileEventsRequest
	GetOperationCode() *string
}

type OperateWebLockFileEventsRequest struct {
	// Specifies whether to handle all alert events. Valid values:
	//
	// - **1**: yes
	//
	// - **0**: no.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DealAll *int32 `json:"DealAll,omitempty" xml:"DealAll,omitempty"`
	// The list of alert event IDs.
	//
	// This parameter is required.
	EventIds []*int64 `json:"EventIds,omitempty" xml:"EventIds,omitempty" type:"Repeated"`
	// The method to handle the alert event. Valid values:
	//
	// - **mark_mis_info**: marks the alert as a false positive
	//
	// - **rm_mark_mis_info**: unmarks the false positive
	//
	// - **offline_handled**: handled offline
	//
	// - **whitelist**: adds to the whitelist
	//
	// - **rm_whitelist**: removes from the whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// whitelist
	OperationCode *string `json:"OperationCode,omitempty" xml:"OperationCode,omitempty"`
}

func (s OperateWebLockFileEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateWebLockFileEventsRequest) GoString() string {
	return s.String()
}

func (s *OperateWebLockFileEventsRequest) GetDealAll() *int32 {
	return s.DealAll
}

func (s *OperateWebLockFileEventsRequest) GetEventIds() []*int64 {
	return s.EventIds
}

func (s *OperateWebLockFileEventsRequest) GetOperationCode() *string {
	return s.OperationCode
}

func (s *OperateWebLockFileEventsRequest) SetDealAll(v int32) *OperateWebLockFileEventsRequest {
	s.DealAll = &v
	return s
}

func (s *OperateWebLockFileEventsRequest) SetEventIds(v []*int64) *OperateWebLockFileEventsRequest {
	s.EventIds = v
	return s
}

func (s *OperateWebLockFileEventsRequest) SetOperationCode(v string) *OperateWebLockFileEventsRequest {
	s.OperationCode = &v
	return s
}

func (s *OperateWebLockFileEventsRequest) Validate() error {
	return dara.Validate(s)
}
