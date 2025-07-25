// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *AddDomainGroupRequest
	GetGroupName() *string
	SetLang(v string) *AddDomainGroupRequest
	GetLang() *string
}

type AddDomainGroupRequest struct {
	// The name of the domain name group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyGroup
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

func (s AddDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *AddDomainGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AddDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDomainGroupRequest) SetGroupName(v string) *AddDomainGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AddDomainGroupRequest) SetLang(v string) *AddDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
