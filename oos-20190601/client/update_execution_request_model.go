// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateExecutionRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateExecutionRequest
	GetDescription() *string
	SetExecutionId(v string) *UpdateExecutionRequest
	GetExecutionId() *string
	SetParameters(v string) *UpdateExecutionRequest
	GetParameters() *string
	SetRegionId(v string) *UpdateExecutionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateExecutionRequest
	GetResourceGroupId() *string
	SetTags(v string) *UpdateExecutionRequest
	GetTags() *string
}

type UpdateExecutionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// Execution description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-c223xxxxxxxxxxxxxxxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The information about the parameters.
	//
	// example:
	//
	// {"Status":"Running"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the execution.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateExecutionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateExecutionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExecutionRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *UpdateExecutionRequest) GetParameters() *string {
	return s.Parameters
}

func (s *UpdateExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateExecutionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateExecutionRequest) GetTags() *string {
	return s.Tags
}

func (s *UpdateExecutionRequest) SetClientToken(v string) *UpdateExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExecutionRequest) SetDescription(v string) *UpdateExecutionRequest {
	s.Description = &v
	return s
}

func (s *UpdateExecutionRequest) SetExecutionId(v string) *UpdateExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *UpdateExecutionRequest) SetParameters(v string) *UpdateExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateExecutionRequest) SetRegionId(v string) *UpdateExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateExecutionRequest) SetResourceGroupId(v string) *UpdateExecutionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateExecutionRequest) SetTags(v string) *UpdateExecutionRequest {
	s.Tags = &v
	return s
}

func (s *UpdateExecutionRequest) Validate() error {
	return dara.Validate(s)
}
