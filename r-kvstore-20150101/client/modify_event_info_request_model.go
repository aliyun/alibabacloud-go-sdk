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
	// The JSON-formatted parameters related to the action. Set this parameter to `{"recoverMode": "xxx", "recoverTime": "xxx"}` if the **TaskAction*	- parameter is set to **modifySwitchTime**.
	//
	// 	- **recoverMode**: specifies the restoration mode for the task. Valid values:
	//
	//     	- **timePoint**: performs the task at the specified point in time.
	//
	//     	- **immediate**: performs the task immediately.
	//
	//     	- **maintainTime**: performs the task within the maintenance window.
	//
	// 	- **recoverTime**: specifies the point in time for restoration. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. This parameter is required if the **recoverMode*	- parameter is set to **timePoint**.
	//
	// example:
	//
	// {"recoverTime":"2023-04-17T14:02:35Z","recoverMode":"timePoint"}
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	// The event handling action. Valid values:
	//
	// 	- **archive**
	//
	// 	- **undo**
	//
	// example:
	//
	// archive
	EventAction *string `json:"EventAction,omitempty" xml:"EventAction,omitempty"`
	// The event IDs. Separate multiple event IDs with commas (,). You can specify up to 20 event IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5422964
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The region ID.
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
