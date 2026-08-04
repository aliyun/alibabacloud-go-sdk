// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v string) *OperateNodeRequest
	GetOperation() *string
	SetOperationParameters(v *NodeOperationParameters) *OperateNodeRequest
	GetOperationParameters() *NodeOperationParameters
	SetResourceGroupId(v string) *OperateNodeRequest
	GetResourceGroupId() *string
}

type OperateNodeRequest struct {
	// The operation to perform.
	//
	// - Cordon: Disables scheduling on the node.
	//
	// - Uncordon: Enables scheduling on the node.
	//
	// - Drain: Evicts pods from the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cordon
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The parameter settings for the node operation.
	OperationParameters *NodeOperationParameters `json:"OperationParameters,omitempty" xml:"OperationParameters,omitempty"`
	// The ID of the resource group that contains the node.
	//
	// example:
	//
	// rgaekzbrflewn6jjq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s OperateNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateNodeRequest) GoString() string {
	return s.String()
}

func (s *OperateNodeRequest) GetOperation() *string {
	return s.Operation
}

func (s *OperateNodeRequest) GetOperationParameters() *NodeOperationParameters {
	return s.OperationParameters
}

func (s *OperateNodeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *OperateNodeRequest) SetOperation(v string) *OperateNodeRequest {
	s.Operation = &v
	return s
}

func (s *OperateNodeRequest) SetOperationParameters(v *NodeOperationParameters) *OperateNodeRequest {
	s.OperationParameters = v
	return s
}

func (s *OperateNodeRequest) SetResourceGroupId(v string) *OperateNodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *OperateNodeRequest) Validate() error {
	if s.OperationParameters != nil {
		if err := s.OperationParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
