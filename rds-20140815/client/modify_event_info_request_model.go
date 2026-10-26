// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionParams(v string) *ModifyEventInfoRequest
	GetActionParams() *string
	SetEventAction(v string) *ModifyEventInfoRequest
	GetEventAction() *string
	SetEventId(v string) *ModifyEventInfoRequest
	GetEventId() *string
	SetRegionId(v string) *ModifyEventInfoRequest
	GetRegionId() *string
	SetSecurityToken(v string) *ModifyEventInfoRequest
	GetSecurityToken() *string
}

type ModifyEventInfoRequest struct {
	// The parameters related to the action. You can add extensions based on your business requirements. When the taskAction parameter is set to modifySwitchTime, set the ActionParams parameter to `{"recoverMode": "xxx", "recoverTime": "xxx"}`.
	//
	// The recoverMode parameter specifies the task recovery pattern. Valid values:
	//
	// - **timePoint**: Executes the task at a specified point in time.
	//
	// - **immediate**: Executes the task immediately.
	//
	// The recoverTime parameter specifies the recovery time. The time must be in UTC+0 and in the yyyy-MM-ddTHH:mm:ssZ format. When the recoverMode parameter is set to timePoint, the recoverTime parameter is required.
	//
	// example:
	//
	// {"recoverTime":"2023-04-17T14:02:35Z","recoverMode":"timePoint"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// The action to process the event. Valid values:
	//
	// - **archive**: Archives the event.
	//
	// - **undo**: Does not process the event.
	//
	// > This parameter is required.
	//
	// example:
	//
	// archive
	EventAction *string `json:"EventAction,omitempty" xml:"EventAction,omitempty"`
	// The event IDs. You can call the DescribeEvents operation to query the event IDs. Separate multiple event IDs with commas (,). You can specify up to 20 event IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5422964
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyEventInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventInfoRequest) GetActionParams() *string {
	return s.ActionParams
}

func (s *ModifyEventInfoRequest) GetEventAction() *string {
	return s.EventAction
}

func (s *ModifyEventInfoRequest) GetEventId() *string {
	return s.EventId
}

func (s *ModifyEventInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEventInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyEventInfoRequest) SetActionParams(v string) *ModifyEventInfoRequest {
	s.ActionParams = &v
	return s
}

func (s *ModifyEventInfoRequest) SetEventAction(v string) *ModifyEventInfoRequest {
	s.EventAction = &v
	return s
}

func (s *ModifyEventInfoRequest) SetEventId(v string) *ModifyEventInfoRequest {
	s.EventId = &v
	return s
}

func (s *ModifyEventInfoRequest) SetRegionId(v string) *ModifyEventInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEventInfoRequest) SetSecurityToken(v string) *ModifyEventInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyEventInfoRequest) Validate() error {
	return dara.Validate(s)
}
