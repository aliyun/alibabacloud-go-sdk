// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddDomainRequest
	GetDomainName() *string
	SetGroupId(v string) *AddDomainRequest
	GetGroupId() *string
	SetLang(v string) *AddDomainRequest
	GetLang() *string
	SetResourceGroupId(v string) *AddDomainRequest
	GetResourceGroupId() *string
}

type AddDomainRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the domain name group. If you do not specify this parameter, pass an empty string, or pass defaultGroup, the domain name is added to the default group. You can call the AddDomainGroup operation to obtain the group ID.
	//
	// example:
	//
	// defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the resource group in Resource Management to which the domain name belongs. If you leave this parameter empty, the domain name is added to the default resource group.
	//
	// example:
	//
	// ******resourcegroupid
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AddDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AddDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddDomainRequest) SetDomainName(v string) *AddDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainRequest) SetGroupId(v string) *AddDomainRequest {
	s.GroupId = &v
	return s
}

func (s *AddDomainRequest) SetLang(v string) *AddDomainRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainRequest) SetResourceGroupId(v string) *AddDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddDomainRequest) Validate() error {
	return dara.Validate(s)
}
