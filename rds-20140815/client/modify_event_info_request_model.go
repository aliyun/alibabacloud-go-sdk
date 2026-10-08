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
	// The action-related parameters, which can be an extension based on business requirements. When taskAction is set to modifySwitchTime, set ActionParams to `{"recoverMode": "xxx", "recoverTime": "xxx"}`.
	//
	// recoverMode specifies the task recovery pattern. Valid values:
	//
	// - **timePoint**: Executes at a specified point in time.
	//
	// - **immediate**: Executes immediately.
	//
	// recoverTime specifies the recovery time in UTC+0. Format: yyyy-MM-ddTHH:mm:ssZ. This parameter is required when recoverMode is set to timePoint.
	//
	// example:
	//
	// {"recoverTime":"2023-04-17T14:02:35Z","recoverMode":"timePoint"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// The event action. Valid values:
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
	// The event ID. You can call the DescribeEvents operation to query event IDs. To query multiple events, separate the event IDs with commas (,). A maximum of 20 event IDs are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5422964
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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
