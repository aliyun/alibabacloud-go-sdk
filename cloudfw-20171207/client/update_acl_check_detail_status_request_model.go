// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclCheckDetailStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateAclCheckDetailStatusRequest
	GetLang() *string
	SetStatus(v string) *UpdateAclCheckDetailStatusRequest
	GetStatus() *string
	SetTaskId(v string) *UpdateAclCheckDetailStatusRequest
	GetTaskId() *string
	SetUuid(v string) *UpdateAclCheckDetailStatusRequest
	GetUuid() *string
}

type UpdateAclCheckDetailStatusRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// task-c92d4544ef7b6a42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bbbb43c9-a931-4d89-9939-86d509139a20
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UpdateAclCheckDetailStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclCheckDetailStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAclCheckDetailStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateAclCheckDetailStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAclCheckDetailStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAclCheckDetailStatusRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateAclCheckDetailStatusRequest) SetLang(v string) *UpdateAclCheckDetailStatusRequest {
	s.Lang = &v
	return s
}

func (s *UpdateAclCheckDetailStatusRequest) SetStatus(v string) *UpdateAclCheckDetailStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateAclCheckDetailStatusRequest) SetTaskId(v string) *UpdateAclCheckDetailStatusRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAclCheckDetailStatusRequest) SetUuid(v string) *UpdateAclCheckDetailStatusRequest {
	s.Uuid = &v
	return s
}

func (s *UpdateAclCheckDetailStatusRequest) Validate() error {
	return dara.Validate(s)
}
