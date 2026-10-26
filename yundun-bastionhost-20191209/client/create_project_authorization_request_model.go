// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProjectAuthorizationRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateProjectAuthorizationRequest
	GetInstanceId() *string
	SetName(v string) *CreateProjectAuthorizationRequest
	GetName() *string
	SetRegionId(v string) *CreateProjectAuthorizationRequest
	GetRegionId() *string
	SetScopeType(v int32) *CreateProjectAuthorizationRequest
	GetScopeType() *int32
}

type CreateProjectAuthorizationRequest struct {
	// The description of the authorization. The description can be up to 500 characters in length.
	//
	// example:
	//
	// Bastionhost demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-5yd34ol020a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The authorization name. The name must be 1 to 128 characters in length. It cannot start with a special character and can contain only the following special characters: periods (.), underscores (_), hyphens (-), backslashes (\\), and spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of manageable assets associated with the authorization. Valid values:
	//
	// - **1**: All projects.
	//
	// - **2**: Specified projects.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ScopeType *int32 `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s CreateProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectAuthorizationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateProjectAuthorizationRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProjectAuthorizationRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *CreateProjectAuthorizationRequest) SetDescription(v string) *CreateProjectAuthorizationRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectAuthorizationRequest) SetInstanceId(v string) *CreateProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateProjectAuthorizationRequest) SetName(v string) *CreateProjectAuthorizationRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectAuthorizationRequest) SetRegionId(v string) *CreateProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProjectAuthorizationRequest) SetScopeType(v int32) *CreateProjectAuthorizationRequest {
	s.ScopeType = &v
	return s
}

func (s *CreateProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
