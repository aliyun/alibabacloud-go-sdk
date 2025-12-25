// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourcePackStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetSourcePackStatusRequest
	GetTaskId() *string
}

type GetSourcePackStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hjsyuyiuwe7wehg****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetSourcePackStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourcePackStatusRequest) GoString() string {
	return s.String()
}

func (s *GetSourcePackStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetSourcePackStatusRequest) SetTaskId(v string) *GetSourcePackStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetSourcePackStatusRequest) Validate() error {
	return dara.Validate(s)
}
