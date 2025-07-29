// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteWorkflowResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWorkflowResponseBody
	GetSuccess() *bool
}

type DeleteWorkflowResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// Your request is denied as lack of ssl protect.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the workflow was deleted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWorkflowResponseBody) SetCode(v int32) *DeleteWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkflowResponseBody) SetMessage(v string) *DeleteWorkflowResponseBody {
	s.Message = &v
	return s
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
