// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindPkByHidForLoginWithLegacyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHid(v string) *FindPkByHidForLoginWithLegacyRequest
	GetHid() *string
	SetSecurityToken(v string) *FindPkByHidForLoginWithLegacyRequest
	GetSecurityToken() *string
}

type FindPkByHidForLoginWithLegacyRequest struct {
	Hid           *string `json:"Hid,omitempty" xml:"Hid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s FindPkByHidForLoginWithLegacyRequest) String() string {
	return dara.Prettify(s)
}

func (s FindPkByHidForLoginWithLegacyRequest) GoString() string {
	return s.String()
}

func (s *FindPkByHidForLoginWithLegacyRequest) GetHid() *string {
	return s.Hid
}

func (s *FindPkByHidForLoginWithLegacyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *FindPkByHidForLoginWithLegacyRequest) SetHid(v string) *FindPkByHidForLoginWithLegacyRequest {
	s.Hid = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyRequest) SetSecurityToken(v string) *FindPkByHidForLoginWithLegacyRequest {
	s.SecurityToken = &v
	return s
}

func (s *FindPkByHidForLoginWithLegacyRequest) Validate() error {
	return dara.Validate(s)
}
