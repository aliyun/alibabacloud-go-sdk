// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPackSceneTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetPackSceneTaskStatusRequest
	GetTaskId() *string
	SetType(v string) *GetPackSceneTaskStatusRequest
	GetType() *string
}

type GetPackSceneTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hjsyuyiuwe7wehg****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// download
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPackSceneTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPackSceneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPackSceneTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetPackSceneTaskStatusRequest) GetType() *string {
	return s.Type
}

func (s *GetPackSceneTaskStatusRequest) SetTaskId(v string) *GetPackSceneTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetPackSceneTaskStatusRequest) SetType(v string) *GetPackSceneTaskStatusRequest {
	s.Type = &v
	return s
}

func (s *GetPackSceneTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
