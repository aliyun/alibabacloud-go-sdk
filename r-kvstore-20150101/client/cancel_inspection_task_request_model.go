// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CancelInspectionTaskRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *CancelInspectionTaskRequest
	GetSecurityToken() *string
	SetTaskId(v string) *CancelInspectionTaskRequest
	GetTaskId() *string
}

type CancelInspectionTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ta-bp1zxszhcgatnx****
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelInspectionTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelInspectionTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CancelInspectionTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelInspectionTaskRequest) SetInstanceId(v string) *CancelInspectionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelInspectionTaskRequest) SetSecurityToken(v string) *CancelInspectionTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *CancelInspectionTaskRequest) SetTaskId(v string) *CancelInspectionTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelInspectionTaskRequest) Validate() error {
	return dara.Validate(s)
}
