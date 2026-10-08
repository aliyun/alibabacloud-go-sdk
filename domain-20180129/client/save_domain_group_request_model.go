// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroupId(v int64) *SaveDomainGroupRequest
	GetDomainGroupId() *int64
	SetDomainGroupName(v string) *SaveDomainGroupRequest
	GetDomainGroupName() *string
	SetLang(v string) *SaveDomainGroupRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveDomainGroupRequest
	GetUserClientIp() *string
}

type SaveDomainGroupRequest struct {
	// Domain group ID. If this parameter is not provided, a new group is created. If it is provided, the domain group name is updated.
	//
	// example:
	//
	// 123456
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// Domain Name Group Name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试分组
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// Language for error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese;
	//
	// - **en**: English.
	//
	// Default value is **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// User IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *SaveDomainGroupRequest) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *SaveDomainGroupRequest) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *SaveDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveDomainGroupRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveDomainGroupRequest) SetDomainGroupId(v int64) *SaveDomainGroupRequest {
	s.DomainGroupId = &v
	return s
}

func (s *SaveDomainGroupRequest) SetDomainGroupName(v string) *SaveDomainGroupRequest {
	s.DomainGroupName = &v
	return s
}

func (s *SaveDomainGroupRequest) SetLang(v string) *SaveDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *SaveDomainGroupRequest) SetUserClientIp(v string) *SaveDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
