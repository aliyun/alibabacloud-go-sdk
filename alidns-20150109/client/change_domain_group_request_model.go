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
	// The domain name.<props="china"> Call [DescribeDomains](https://help.aliyun.com/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c4g.11186623.help-menu-search-29697.d_0) to obtain the domain name.
	//
	// <props="intl">Call [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the target domain name group.
	//
	// - If you do not specify GroupId, the domain name is moved to the default group.
	//
	// - If GroupId is an empty string (""), the domain name is moved to the default group.
	//
	// - If GroupId is defaultGroup, the domain name is moved to the default group.
	//
	// - If GroupId is a different value, the system checks if the group exists. If the group exists, the domain name\\"s group is updated. If the group does not exist, the group is not updated.
	//
	// example:
	//
	// 60bb3ef15ace449082cf914ed3ea****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese
	//
	// - en: English
	//
	// Default: en.
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
