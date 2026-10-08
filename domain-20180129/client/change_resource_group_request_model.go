// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ChangeResourceGroupRequest
	GetLang() *string
	SetNewResourceGroupId(v string) *ChangeResourceGroupRequest
	GetNewResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
	SetUserClientIp(v string) *ChangeResourceGroupRequest
	GetUserClientIp() *string
}

type ChangeResourceGroupRequest struct {
	// The language in which error messages are returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **zh**.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the resource group to which you want to shift the domain name.
	//
	// You can view the resource group ID in the [Resource Management Console](https://resourcemanager.console.aliyun.com/resource-groups).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek2tcx7os7bkmq
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The resource ID of the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// S20227H17A561968
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type of the domain name. This parameter is fixed to “Domain” and does not need to be specified.
	//
	// example:
	//
	// Domain
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The IP address of the user client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *ChangeResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ChangeResourceGroupRequest) SetLang(v string) *ChangeResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
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

func (s *ChangeResourceGroupRequest) SetUserClientIp(v string) *ChangeResourceGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
