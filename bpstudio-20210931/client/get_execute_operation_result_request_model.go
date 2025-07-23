// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteOperationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *GetExecuteOperationResultRequest
	GetOperationId() *string
	SetResourceGroupId(v string) *GetExecuteOperationResultRequest
	GetResourceGroupId() *string
}

type GetExecuteOperationResultRequest struct {
	// The ID of the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// op_xxxxxxxxxxxxxxxxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the resource group. This parameter is specified to verify the permissions on the resource group.
	//
	// example:
	//
	// testResourceId
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetExecuteOperationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteOperationResultRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *GetExecuteOperationResultRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetExecuteOperationResultRequest) SetOperationId(v string) *GetExecuteOperationResultRequest {
	s.OperationId = &v
	return s
}

func (s *GetExecuteOperationResultRequest) SetResourceGroupId(v string) *GetExecuteOperationResultRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetExecuteOperationResultRequest) Validate() error {
	return dara.Validate(s)
}
