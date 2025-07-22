// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEndpointSwitchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetEndpointSwitchTaskRequest
	GetTaskId() *string
	SetUid(v string) *GetEndpointSwitchTaskRequest
	GetUid() *string
	SetUserId(v string) *GetEndpointSwitchTaskRequest
	GetUserId() *string
	SetContext(v string) *GetEndpointSwitchTaskRequest
	GetContext() *string
	SetAccessKey(v string) *GetEndpointSwitchTaskRequest
	GetAccessKey() *string
	SetSignature(v string) *GetEndpointSwitchTaskRequest
	GetSignature() *string
	SetSkipAuth(v string) *GetEndpointSwitchTaskRequest
	GetSkipAuth() *string
	SetTimestamp(v string) *GetEndpointSwitchTaskRequest
	GetTimestamp() *string
}

type GetEndpointSwitchTaskRequest struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Context   *string `json:"__context,omitempty" xml:"__context,omitempty"`
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	SkipAuth  *string `json:"skipAuth,omitempty" xml:"skipAuth,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s GetEndpointSwitchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointSwitchTaskRequest) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEndpointSwitchTaskRequest) GetUid() *string {
	return s.Uid
}

func (s *GetEndpointSwitchTaskRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetEndpointSwitchTaskRequest) GetContext() *string {
	return s.Context
}

func (s *GetEndpointSwitchTaskRequest) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetEndpointSwitchTaskRequest) GetSignature() *string {
	return s.Signature
}

func (s *GetEndpointSwitchTaskRequest) GetSkipAuth() *string {
	return s.SkipAuth
}

func (s *GetEndpointSwitchTaskRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetEndpointSwitchTaskRequest) SetTaskId(v string) *GetEndpointSwitchTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetUid(v string) *GetEndpointSwitchTaskRequest {
	s.Uid = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetUserId(v string) *GetEndpointSwitchTaskRequest {
	s.UserId = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetContext(v string) *GetEndpointSwitchTaskRequest {
	s.Context = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetAccessKey(v string) *GetEndpointSwitchTaskRequest {
	s.AccessKey = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetSignature(v string) *GetEndpointSwitchTaskRequest {
	s.Signature = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetSkipAuth(v string) *GetEndpointSwitchTaskRequest {
	s.SkipAuth = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) SetTimestamp(v string) *GetEndpointSwitchTaskRequest {
	s.Timestamp = &v
	return s
}

func (s *GetEndpointSwitchTaskRequest) Validate() error {
	return dara.Validate(s)
}
