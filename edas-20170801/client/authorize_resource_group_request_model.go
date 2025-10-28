// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupIds(v string) *AuthorizeResourceGroupRequest
	GetResourceGroupIds() *string
	SetTargetUserId(v string) *AuthorizeResourceGroupRequest
	GetTargetUserId() *string
}

type AuthorizeResourceGroupRequest struct {
	// The ID of the resource group. You can call the ListResourceGroup operation to query the resource group ID. For more information, see [ListResourceGroup](https://help.aliyun.com/document_detail/62055.html).
	//
	// You can specify multiple resource group IDs. Separate multiple resource group IDs with semicolons (;).
	//
	// This parameter is required.
	//
	// example:
	//
	// 461;462
	ResourceGroupIds *string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty"`
	// The ID of the RAM user to be authorized.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@13333********
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s AuthorizeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeResourceGroupRequest) GetResourceGroupIds() *string {
	return s.ResourceGroupIds
}

func (s *AuthorizeResourceGroupRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *AuthorizeResourceGroupRequest) SetResourceGroupIds(v string) *AuthorizeResourceGroupRequest {
	s.ResourceGroupIds = &v
	return s
}

func (s *AuthorizeResourceGroupRequest) SetTargetUserId(v string) *AuthorizeResourceGroupRequest {
	s.TargetUserId = &v
	return s
}

func (s *AuthorizeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
