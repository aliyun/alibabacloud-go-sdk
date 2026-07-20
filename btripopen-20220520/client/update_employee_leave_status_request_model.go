// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeLeaveStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsLeave(v bool) *UpdateEmployeeLeaveStatusRequest
	GetIsLeave() *bool
	SetUserId(v string) *UpdateEmployeeLeaveStatusRequest
	GetUserId() *string
}

type UpdateEmployeeLeaveStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	IsLeave *bool `json:"is_leave,omitempty" xml:"is_leave,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s UpdateEmployeeLeaveStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeLeaveStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeLeaveStatusRequest) GetIsLeave() *bool {
	return s.IsLeave
}

func (s *UpdateEmployeeLeaveStatusRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateEmployeeLeaveStatusRequest) SetIsLeave(v bool) *UpdateEmployeeLeaveStatusRequest {
	s.IsLeave = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusRequest) SetUserId(v string) *UpdateEmployeeLeaveStatusRequest {
	s.UserId = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusRequest) Validate() error {
	return dara.Validate(s)
}
