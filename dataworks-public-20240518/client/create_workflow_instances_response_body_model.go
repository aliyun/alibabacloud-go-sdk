// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *CreateWorkflowInstancesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *CreateWorkflowInstancesResponseBody
	GetRequestId() *string
}

type CreateWorkflowInstancesResponseBody struct {
	// The operation ID. You can use this value to call the GetCreateWorkflowInstancesResult operation to query the result of the creation operation.
	//
	// example:
	//
	// e15ad21c-b0e9-4792-8f55-b037xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *CreateWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowInstancesResponseBody) SetOperationId(v string) *CreateWorkflowInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateWorkflowInstancesResponseBody) SetRequestId(v string) *CreateWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
