// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInspectionScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteInspectionScheduleRequest
	GetInstanceId() *string
	SetScheduleId(v string) *DeleteInspectionScheduleRequest
	GetScheduleId() *string
	SetSecurityToken(v string) *DeleteInspectionScheduleRequest
	GetSecurityToken() *string
}

type DeleteInspectionScheduleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ta-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// scheduleId-202604141xxxx
	ScheduleId    *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteInspectionScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInspectionScheduleRequest) GoString() string {
	return s.String()
}

func (s *DeleteInspectionScheduleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInspectionScheduleRequest) GetScheduleId() *string {
	return s.ScheduleId
}

func (s *DeleteInspectionScheduleRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteInspectionScheduleRequest) SetInstanceId(v string) *DeleteInspectionScheduleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInspectionScheduleRequest) SetScheduleId(v string) *DeleteInspectionScheduleRequest {
	s.ScheduleId = &v
	return s
}

func (s *DeleteInspectionScheduleRequest) SetSecurityToken(v string) *DeleteInspectionScheduleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteInspectionScheduleRequest) Validate() error {
	return dara.Validate(s)
}
