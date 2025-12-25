// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDisableWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateDisableWorkflowsResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateDisableWorkflowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateDisableWorkflowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateDisableWorkflowsResponseBody
	GetSuccess() *bool
}

type OperateDisableWorkflowsResponseBody struct {
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
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateDisableWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateDisableWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateDisableWorkflowsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateDisableWorkflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateDisableWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateDisableWorkflowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateDisableWorkflowsResponseBody) SetCode(v int32) *OperateDisableWorkflowsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateDisableWorkflowsResponseBody) SetMessage(v string) *OperateDisableWorkflowsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateDisableWorkflowsResponseBody) SetRequestId(v string) *OperateDisableWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateDisableWorkflowsResponseBody) SetSuccess(v bool) *OperateDisableWorkflowsResponseBody {
	s.Success = &v
	return s
}

func (s *OperateDisableWorkflowsResponseBody) Validate() error {
	return dara.Validate(s)
}
