// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRerunWorkflowInstancesResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *GetRerunWorkflowInstancesResultRequest
	GetOperationId() *string
}

type GetRerunWorkflowInstancesResultRequest struct {
	// The operation ID used to asynchronously query the result of the workflow instance rerun. This value is obtained from the RerunWorkflowInstances operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// e15ad21c-b0e9-4792-8f55-b037xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s GetRerunWorkflowInstancesResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRerunWorkflowInstancesResultRequest) GoString() string {
	return s.String()
}

func (s *GetRerunWorkflowInstancesResultRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *GetRerunWorkflowInstancesResultRequest) SetOperationId(v string) *GetRerunWorkflowInstancesResultRequest {
	s.OperationId = &v
	return s
}

func (s *GetRerunWorkflowInstancesResultRequest) Validate() error {
	return dara.Validate(s)
}
