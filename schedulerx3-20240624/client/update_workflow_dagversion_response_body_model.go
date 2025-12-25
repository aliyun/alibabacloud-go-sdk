// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowDAGVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateWorkflowDAGVersionResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateWorkflowDAGVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWorkflowDAGVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkflowDAGVersionResponseBody
	GetSuccess() *bool
}

type UpdateWorkflowDAGVersionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C78E2AD2-5985-515B-BAD2-31A248AFC263
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDAGVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowDAGVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDAGVersionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateWorkflowDAGVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowDAGVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowDAGVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkflowDAGVersionResponseBody) SetCode(v int32) *UpdateWorkflowDAGVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowDAGVersionResponseBody) SetMessage(v string) *UpdateWorkflowDAGVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowDAGVersionResponseBody) SetRequestId(v string) *UpdateWorkflowDAGVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDAGVersionResponseBody) SetSuccess(v bool) *UpdateWorkflowDAGVersionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkflowDAGVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
