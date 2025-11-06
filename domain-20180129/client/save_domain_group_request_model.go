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
	// example:
	//
	// 123456
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// This parameter is required.
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
