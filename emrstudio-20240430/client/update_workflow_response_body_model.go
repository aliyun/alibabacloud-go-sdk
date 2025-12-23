// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkflowResponseBody
	GetSuccess() *bool
}

type UpdateWorkflowResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkflowResponseBody) SetRequestId(v string) *UpdateWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowResponseBody) SetSuccess(v bool) *UpdateWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
