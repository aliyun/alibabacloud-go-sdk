// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *RerunWorkflowInstancesResponseBody
	GetOperationId() *string
	SetRequestId(v string) *RerunWorkflowInstancesResponseBody
	GetRequestId() *string
}

type RerunWorkflowInstancesResponseBody struct {
	// The operation ID. You can use this value to query the creation result via the GetRerunWorkflowInstancesResult operation.
	//
	// example:
	//
	// e15ad21c-b0e9-4792-8f55-b037xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID. Used for troubleshooting and log tracking.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RerunWorkflowInstancesResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *RerunWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RerunWorkflowInstancesResponseBody) SetOperationId(v string) *RerunWorkflowInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *RerunWorkflowInstancesResponseBody) SetRequestId(v string) *RerunWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerunWorkflowInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
