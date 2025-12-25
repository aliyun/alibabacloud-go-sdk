// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateEnableWorkflowsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateEnableWorkflowsResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateEnableWorkflowsResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateEnableWorkflowsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateEnableWorkflowsResponseBody
	GetSuccess() *bool
}

type OperateEnableWorkflowsResponseBody struct {
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
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateEnableWorkflowsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateEnableWorkflowsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateEnableWorkflowsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateEnableWorkflowsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateEnableWorkflowsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateEnableWorkflowsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateEnableWorkflowsResponseBody) SetCode(v int32) *OperateEnableWorkflowsResponseBody {
	s.Code = &v
	return s
}

func (s *OperateEnableWorkflowsResponseBody) SetMessage(v string) *OperateEnableWorkflowsResponseBody {
	s.Message = &v
	return s
}

func (s *OperateEnableWorkflowsResponseBody) SetRequestId(v string) *OperateEnableWorkflowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateEnableWorkflowsResponseBody) SetSuccess(v bool) *OperateEnableWorkflowsResponseBody {
	s.Success = &v
	return s
}

func (s *OperateEnableWorkflowsResponseBody) Validate() error {
	return dara.Validate(s)
}
