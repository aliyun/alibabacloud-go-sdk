// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ChangeDomainGroupRequest
	GetDomainName() *string
	SetGroupId(v string) *ChangeDomainGroupRequest
	GetGroupId() *string
	SetLang(v string) *ChangeDomainGroupRequest
	GetLang() *string
}

type ChangeDomainGroupRequest struct {
	// The domain name. You can call the [DescribeDomains ](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0)operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the target domain name group.
	//
	// 	- If you do not specify GroupId, the domain name belongs to the default group.
	//
	// 	- If you specify an empty string "" for GroupId, the domain name belongs to the default group.
	//
	// 	- If you set GroupId to defaultGroup, the domain name belongs to the default group.
	//
	// 	- If you specify GroupId to another value and the value is verified existent, the value of GroupId for the target domain name is updated. If the value is verified inexistent, the value of GroupId for the target domain name is not updated.
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
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ChangeDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeDomainGroupRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ChangeDomainGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ChangeDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *ChangeDomainGroupRequest) SetDomainName(v string) *ChangeDomainGroupRequest {
	s.DomainName = &v
	return s
}

func (s *ChangeDomainGroupRequest) SetGroupId(v string) *ChangeDomainGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ChangeDomainGroupRequest) SetLang(v string) *ChangeDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *ChangeDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
