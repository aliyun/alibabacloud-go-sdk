// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveDomainResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *MoveDomainResourceGroupRequest
	GetLang() *string
	SetNewResourceGroupId(v string) *MoveDomainResourceGroupRequest
	GetNewResourceGroupId() *string
	SetResourceId(v string) *MoveDomainResourceGroupRequest
	GetResourceId() *string
}

type MoveDomainResourceGroupRequest struct {
	// The language of the values of specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the new resource group.
	//
	// You can view the resource group ID in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups?).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzzk7hx3glaoq
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s MoveDomainResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveDomainResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveDomainResourceGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *MoveDomainResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *MoveDomainResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveDomainResourceGroupRequest) SetLang(v string) *MoveDomainResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveDomainResourceGroupRequest) SetNewResourceGroupId(v string) *MoveDomainResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveDomainResourceGroupRequest) SetResourceId(v string) *MoveDomainResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveDomainResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
