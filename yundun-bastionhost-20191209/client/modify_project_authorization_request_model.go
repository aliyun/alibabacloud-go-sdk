// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *ModifyProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetDescription(v string) *ModifyProjectAuthorizationRequest
	GetDescription() *string
	SetInstanceId(v string) *ModifyProjectAuthorizationRequest
	GetInstanceId() *string
	SetName(v string) *ModifyProjectAuthorizationRequest
	GetName() *string
	SetRegionId(v string) *ModifyProjectAuthorizationRequest
	GetRegionId() *string
	SetScopeType(v int32) *ModifyProjectAuthorizationRequest
	GetScopeType() *int32
}

type ModifyProjectAuthorizationRequest struct {
	// The ID of the project authorization to be modified.
	//
	// > You can call the ListProjectAuthorizations operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The new description of the project authorization.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the project authorization.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scope of assets that can be managed by the authorization after modification. Valid values:
	//
	// - **1**: all projects
	//
	// - **2**: specified projects
	//
	// example:
	//
	// 1
	ScopeType *int32 `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s ModifyProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *ModifyProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *ModifyProjectAuthorizationRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyProjectAuthorizationRequest) GetName() *string {
	return s.Name
}

func (s *ModifyProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyProjectAuthorizationRequest) GetScopeType() *int32 {
	return s.ScopeType
}

func (s *ModifyProjectAuthorizationRequest) SetAuthorizationId(v int64) *ModifyProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *ModifyProjectAuthorizationRequest) SetDescription(v string) *ModifyProjectAuthorizationRequest {
	s.Description = &v
	return s
}

func (s *ModifyProjectAuthorizationRequest) SetInstanceId(v string) *ModifyProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyProjectAuthorizationRequest) SetName(v string) *ModifyProjectAuthorizationRequest {
	s.Name = &v
	return s
}

func (s *ModifyProjectAuthorizationRequest) SetRegionId(v string) *ModifyProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyProjectAuthorizationRequest) SetScopeType(v int32) *ModifyProjectAuthorizationRequest {
	s.ScopeType = &v
	return s
}

func (s *ModifyProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
