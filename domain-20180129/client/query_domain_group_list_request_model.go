// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroupName(v string) *QueryDomainGroupListRequest
	GetDomainGroupName() *string
	SetLang(v string) *QueryDomainGroupListRequest
	GetLang() *string
	SetOrderByType(v string) *QueryDomainGroupListRequest
	GetOrderByType() *string
	SetOrderKeyType(v string) *QueryDomainGroupListRequest
	GetOrderKeyType() *string
	SetShowDeletingGroup(v bool) *QueryDomainGroupListRequest
	GetShowDeletingGroup() *bool
	SetUserClientIp(v string) *QueryDomainGroupListRequest
	GetUserClientIp() *string
}

type QueryDomainGroupListRequest struct {
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// example:
	//
	// en
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OrderByType  *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	OrderKeyType *string `json:"OrderKeyType,omitempty" xml:"OrderKeyType,omitempty"`
	// example:
	//
	// false
	ShowDeletingGroup *bool `json:"ShowDeletingGroup,omitempty" xml:"ShowDeletingGroup,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainGroupListRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListRequest) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *QueryDomainGroupListRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainGroupListRequest) GetOrderByType() *string {
	return s.OrderByType
}

func (s *QueryDomainGroupListRequest) GetOrderKeyType() *string {
	return s.OrderKeyType
}

func (s *QueryDomainGroupListRequest) GetShowDeletingGroup() *bool {
	return s.ShowDeletingGroup
}

func (s *QueryDomainGroupListRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainGroupListRequest) SetDomainGroupName(v string) *QueryDomainGroupListRequest {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetLang(v string) *QueryDomainGroupListRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetOrderByType(v string) *QueryDomainGroupListRequest {
	s.OrderByType = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetOrderKeyType(v string) *QueryDomainGroupListRequest {
	s.OrderKeyType = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetShowDeletingGroup(v bool) *QueryDomainGroupListRequest {
	s.ShowDeletingGroup = &v
	return s
}

func (s *QueryDomainGroupListRequest) SetUserClientIp(v string) *QueryDomainGroupListRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainGroupListRequest) Validate() error {
	return dara.Validate(s)
}
