// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainGroupId(v int64) *DeleteDomainGroupRequest
	GetDomainGroupId() *int64
	SetLang(v string) *DeleteDomainGroupRequest
	GetLang() *string
	SetUserClientIp(v string) *DeleteDomainGroupRequest
	GetUserClientIp() *string
}

type DeleteDomainGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteDomainGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainGroupRequest) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *DeleteDomainGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDomainGroupRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteDomainGroupRequest) SetDomainGroupId(v int64) *DeleteDomainGroupRequest {
	s.DomainGroupId = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetLang(v string) *DeleteDomainGroupRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDomainGroupRequest) SetUserClientIp(v string) *DeleteDomainGroupRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteDomainGroupRequest) Validate() error {
	return dara.Validate(s)
}
