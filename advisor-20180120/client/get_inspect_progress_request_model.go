// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInspectProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *GetInspectProgressRequest
	GetAssumeAliyunId() *int64
	SetTaskId(v int64) *GetInspectProgressRequest
	GetTaskId() *int64
	SetToken(v string) *GetInspectProgressRequest
	GetToken() *string
}

type GetInspectProgressRequest struct {
	// example:
	//
	// 14********37
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// 95***135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetInspectProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInspectProgressRequest) GoString() string {
	return s.String()
}

func (s *GetInspectProgressRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *GetInspectProgressRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetInspectProgressRequest) GetToken() *string {
	return s.Token
}

func (s *GetInspectProgressRequest) SetAssumeAliyunId(v int64) *GetInspectProgressRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *GetInspectProgressRequest) SetTaskId(v int64) *GetInspectProgressRequest {
	s.TaskId = &v
	return s
}

func (s *GetInspectProgressRequest) SetToken(v string) *GetInspectProgressRequest {
	s.Token = &v
	return s
}

func (s *GetInspectProgressRequest) Validate() error {
	return dara.Validate(s)
}
