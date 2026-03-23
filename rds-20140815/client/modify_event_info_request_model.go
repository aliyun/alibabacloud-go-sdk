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
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	EventAction  *string `json:"EventAction,omitempty" xml:"EventAction,omitempty"`
	// This parameter is required.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
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
