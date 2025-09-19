// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitOperationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperateCode(v string) *SubmitOperationTaskResponseBody
	GetOperateCode() *string
	SetRequestId(v string) *SubmitOperationTaskResponseBody
	GetRequestId() *string
	SetRootTaskId(v string) *SubmitOperationTaskResponseBody
	GetRootTaskId() *string
}

type SubmitOperationTaskResponseBody struct {
	// The handling result code. Valid values:
	//
	// 	- Insufficient authorization: AuthorizationExhaust
	//
	// 	- Unauthorized: ActionTrialUnauthorized
	//
	// example:
	//
	// AuthorizationExhaust
	OperateCode *string `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0C8487EF-50C2-54BB-8634-10F8C35D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The main task ID that is returned when the task is submitted.
	//
	// example:
	//
	// 89f5d7813bd59dd237580a8664b3xxxx
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
}

func (s SubmitOperationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitOperationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOperationTaskResponseBody) GetOperateCode() *string {
	return s.OperateCode
}

func (s *SubmitOperationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitOperationTaskResponseBody) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *SubmitOperationTaskResponseBody) SetOperateCode(v string) *SubmitOperationTaskResponseBody {
	s.OperateCode = &v
	return s
}

func (s *SubmitOperationTaskResponseBody) SetRequestId(v string) *SubmitOperationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitOperationTaskResponseBody) SetRootTaskId(v string) *SubmitOperationTaskResponseBody {
	s.RootTaskId = &v
	return s
}

func (s *SubmitOperationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
