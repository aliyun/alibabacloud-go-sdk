// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateWorkflowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateWorkflowInstanceResponseBody
	GetSuccess() *bool
}

type OperateWorkflowInstanceResponseBody struct {
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

func (s OperateWorkflowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *OperateWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateWorkflowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateWorkflowInstanceResponseBody) SetRequestId(v string) *OperateWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateWorkflowInstanceResponseBody) SetSuccess(v bool) *OperateWorkflowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *OperateWorkflowInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
