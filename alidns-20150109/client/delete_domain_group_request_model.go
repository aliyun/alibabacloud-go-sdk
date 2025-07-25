// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteDomainGroupRequest
	GetGroupId() *string
	SetLang(v string) *DeleteDomainGroupRequest
	GetLang() *string
}

type DeleteDomainGroupRequest struct {
	// The ID of the domain name group. You can call the [DescribeDomainGroups](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomaingroups?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2223
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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

func (s DeleteDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDomainGroupRequest) SetGroupId(v string) *DeleteDomainGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetLang(v string) *DeleteDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
