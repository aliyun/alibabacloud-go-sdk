// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *UpdateDomainGroupRequest
	GetGroupId() *string
	SetGroupName(v string) *UpdateDomainGroupRequest
	GetGroupName() *string
	SetLang(v string) *UpdateDomainGroupRequest
	GetLang() *string
}

type UpdateDomainGroupRequest struct {
	// The ID of the domain name group whose name you want to modify. You can call the [DescribeDomainGroups ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaingroups?spm=a2c63.p38356.help-menu-search-29697.d_0)operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The new name of the domain name group.
	//
	// This parameter is required.
	//
	// example:
	//
	// NewName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s UpdateDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateDomainGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDomainGroupRequest) SetGroupId(v string) *UpdateDomainGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateDomainGroupRequest) SetGroupName(v string) *UpdateDomainGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateDomainGroupRequest) SetLang(v string) *UpdateDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
