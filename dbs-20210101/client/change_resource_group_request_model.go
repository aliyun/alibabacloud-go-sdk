// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ChangeResourceGroupRequest
	GetClientToken() *string
	SetNewResourceGroupId(v string) *ChangeResourceGroupRequest
	GetNewResourceGroupId() *string
	SetRegionCode(v string) *ChangeResourceGroupRequest
	GetRegionCode() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dbs
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource group to which you want to move the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekz4kee6******
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbs1jyajqk******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Set the value to backupplan.
	//
	// This parameter is required.
	//
	// example:
	//
	// backupplan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ChangeResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) SetClientToken(v string) *ChangeResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionCode(v string) *ChangeResourceGroupRequest {
	s.RegionCode = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
