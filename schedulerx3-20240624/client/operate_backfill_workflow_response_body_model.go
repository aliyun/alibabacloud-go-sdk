// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateBackfillWorkflowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *OperateBackfillWorkflowResponseBody
	GetCode() *int32
	SetMessage(v string) *OperateBackfillWorkflowResponseBody
	GetMessage() *string
	SetRequestId(v string) *OperateBackfillWorkflowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateBackfillWorkflowResponseBody
	GetSuccess() *bool
}

type OperateBackfillWorkflowResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateBackfillWorkflowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateBackfillWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *OperateBackfillWorkflowResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *OperateBackfillWorkflowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OperateBackfillWorkflowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateBackfillWorkflowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateBackfillWorkflowResponseBody) SetCode(v int32) *OperateBackfillWorkflowResponseBody {
	s.Code = &v
	return s
}

func (s *OperateBackfillWorkflowResponseBody) SetMessage(v string) *OperateBackfillWorkflowResponseBody {
	s.Message = &v
	return s
}

func (s *OperateBackfillWorkflowResponseBody) SetRequestId(v string) *OperateBackfillWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateBackfillWorkflowResponseBody) SetSuccess(v bool) *OperateBackfillWorkflowResponseBody {
	s.Success = &v
	return s
}

func (s *OperateBackfillWorkflowResponseBody) Validate() error {
	return dara.Validate(s)
}
