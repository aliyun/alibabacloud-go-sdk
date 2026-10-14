// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetC2paSignResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetC2paSignResultRequest
	GetTaskId() *string
}

type GetC2paSignResultRequest struct {
	// The opaque task ID returned by CreateC2paSignTask, up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7093c29e-f472-4e1e-91c6-a43df7c7382e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetC2paSignResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetC2paSignResultRequest) GoString() string {
	return s.String()
}

func (s *GetC2paSignResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetC2paSignResultRequest) SetTaskId(v string) *GetC2paSignResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetC2paSignResultRequest) Validate() error {
	return dara.Validate(s)
}
