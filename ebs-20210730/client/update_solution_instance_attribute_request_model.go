// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateSolutionInstanceAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateSolutionInstanceAttributeRequest
	GetDescription() *string
	SetName(v string) *UpdateSolutionInstanceAttributeRequest
	GetName() *string
	SetRegionId(v string) *UpdateSolutionInstanceAttributeRequest
	GetRegionId() *string
	SetSolutionInstanceId(v string) *UpdateSolutionInstanceAttributeRequest
	GetSolutionInstanceId() *string
}

type UpdateSolutionInstanceAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// defaultDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// defaultName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the dedicated block storage cluster resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// inst-***
	SolutionInstanceId *string `json:"SolutionInstanceId,omitempty" xml:"SolutionInstanceId,omitempty"`
}

func (s UpdateSolutionInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSolutionInstanceAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSolutionInstanceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSolutionInstanceAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSolutionInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSolutionInstanceAttributeRequest) GetSolutionInstanceId() *string {
	return s.SolutionInstanceId
}

func (s *UpdateSolutionInstanceAttributeRequest) SetClientToken(v string) *UpdateSolutionInstanceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetDescription(v string) *UpdateSolutionInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetName(v string) *UpdateSolutionInstanceAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetRegionId(v string) *UpdateSolutionInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) SetSolutionInstanceId(v string) *UpdateSolutionInstanceAttributeRequest {
	s.SolutionInstanceId = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
