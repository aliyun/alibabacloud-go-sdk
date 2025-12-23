// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkflowResponseBody
	GetSuccess() *bool
}

type DeleteWorkflowResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkflowResponseBody) SetRequestId(v string) *DeleteWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetSuccess(v bool) *DeleteWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
