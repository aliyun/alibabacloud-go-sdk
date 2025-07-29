// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DisableWorkflowResponseBody
	GetCode() *int32
	SetMessage(v string) *DisableWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableWorkflowResponseBody
	GetSuccess() *bool
}

type DisableWorkflowResponseBody struct {
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
	// Indicates whether the workflow was disabled. Valid values:
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

func (s DisableWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisableWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableWorkflowResponseBody) SetCode(v int32) *DisableWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *DisableWorkflowResponseBody) SetMessage(v string) *DisableWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *DisableWorkflowResponseBody) SetRequestId(v string) *DisableWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableWorkflowResponseBody) SetSuccess(v bool) *DisableWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *DisableWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
