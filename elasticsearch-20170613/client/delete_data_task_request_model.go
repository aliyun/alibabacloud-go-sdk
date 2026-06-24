// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDataTaskRequest
	GetClientToken() *string
	SetTaskId(v string) *DeleteDataTaskRequest
	GetTaskId() *string
}

type DeleteDataTaskRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The index migration task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// et_cn_0oyg09o96ib40****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DeleteDataTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDataTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteDataTaskRequest) SetClientToken(v string) *DeleteDataTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDataTaskRequest) SetTaskId(v string) *DeleteDataTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DeleteDataTaskRequest) Validate() error {
	return dara.Validate(s)
}
