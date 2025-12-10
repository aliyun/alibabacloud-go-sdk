// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuntimeId(v int32) *GetJobResultRequest
	GetRuntimeId() *int32
}

type GetJobResultRequest struct {
	// The execution ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RuntimeId *int32 `json:"runtimeId,omitempty" xml:"runtimeId,omitempty"`
}

func (s GetJobResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetJobResultRequest) GetRuntimeId() *int32 {
	return s.RuntimeId
}

func (s *GetJobResultRequest) SetRuntimeId(v int32) *GetJobResultRequest {
	s.RuntimeId = &v
	return s
}

func (s *GetJobResultRequest) Validate() error {
	return dara.Validate(s)
}
