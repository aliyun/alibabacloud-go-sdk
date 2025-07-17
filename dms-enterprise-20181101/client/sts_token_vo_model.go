// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStsTokenVO interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *StsTokenVO
	GetAccessKeyId() *string
	SetAccessKeySecret(v string) *StsTokenVO
	GetAccessKeySecret() *string
	SetExpiration(v string) *StsTokenVO
	GetExpiration() *string
	SetSecurityToken(v string) *StsTokenVO
	GetSecurityToken() *string
}

type StsTokenVO struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Expiration      *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s StsTokenVO) String() string {
	return dara.Prettify(s)
}

func (s StsTokenVO) GoString() string {
	return s.String()
}

func (s *StsTokenVO) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *StsTokenVO) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *StsTokenVO) GetExpiration() *string {
	return s.Expiration
}

func (s *StsTokenVO) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *StsTokenVO) SetAccessKeyId(v string) *StsTokenVO {
	s.AccessKeyId = &v
	return s
}

func (s *StsTokenVO) SetAccessKeySecret(v string) *StsTokenVO {
	s.AccessKeySecret = &v
	return s
}

func (s *StsTokenVO) SetExpiration(v string) *StsTokenVO {
	s.Expiration = &v
	return s
}

func (s *StsTokenVO) SetSecurityToken(v string) *StsTokenVO {
	s.SecurityToken = &v
	return s
}

func (s *StsTokenVO) Validate() error {
	return dara.Validate(s)
}
