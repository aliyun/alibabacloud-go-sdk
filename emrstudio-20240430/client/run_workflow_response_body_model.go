// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunWorkflowResponseBody
	GetSuccess() *bool
}

type RunWorkflowResponseBody struct {
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

func (s RunWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *RunWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunWorkflowResponseBody) SetRequestId(v string) *RunWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunWorkflowResponseBody) SetSuccess(v bool) *RunWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *RunWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
