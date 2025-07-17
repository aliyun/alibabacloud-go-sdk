// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateWorkflowInstancesResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *GetCreateWorkflowInstancesResultRequest
	GetOperationId() *string
}

type GetCreateWorkflowInstancesResultRequest struct {
	// The operation ID. This parameter is used to query the result of asynchronously creating a workflow instance. You can call the CreateWorkflowInstances operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e15ad21c-b0e9-4792-8f55-b037xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s GetCreateWorkflowInstancesResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCreateWorkflowInstancesResultRequest) GoString() string {
	return s.String()
}

func (s *GetCreateWorkflowInstancesResultRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *GetCreateWorkflowInstancesResultRequest) SetOperationId(v string) *GetCreateWorkflowInstancesResultRequest {
	s.OperationId = &v
	return s
}

func (s *GetCreateWorkflowInstancesResultRequest) Validate() error {
	return dara.Validate(s)
}
