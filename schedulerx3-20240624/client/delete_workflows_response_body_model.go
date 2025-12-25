// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteWorkflowsResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteWorkflowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkflowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkflowsResponseBody
	GetSuccess() *bool
}

type DeleteWorkflowsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 27B1345D-5F71-5972-8E4C-AABA6C6232F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteWorkflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkflowsResponseBody) SetCode(v int32) *DeleteWorkflowsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkflowsResponseBody) SetMessage(v string) *DeleteWorkflowsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkflowsResponseBody) SetRequestId(v string) *DeleteWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowsResponseBody) SetSuccess(v bool) *DeleteWorkflowsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWorkflowsResponseBody) Validate() error {
	return dara.Validate(s)
}
