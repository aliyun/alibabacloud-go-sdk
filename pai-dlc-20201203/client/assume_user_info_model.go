// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeUserInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *AssumeUserInfo
	GetAccessKeyId() *string
	SetId(v string) *AssumeUserInfo
	GetId() *string
	SetSecurityToken(v string) *AssumeUserInfo
	GetSecurityToken() *string
	SetType(v string) *AssumeUserInfo
	GetType() *string
}

type AssumeUserInfo struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AssumeUserInfo) String() string {
	return dara.Prettify(s)
}

func (s AssumeUserInfo) GoString() string {
	return s.String()
}

func (s *AssumeUserInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *AssumeUserInfo) GetId() *string {
	return s.Id
}

func (s *AssumeUserInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AssumeUserInfo) GetType() *string {
	return s.Type
}

func (s *AssumeUserInfo) SetAccessKeyId(v string) *AssumeUserInfo {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeUserInfo) SetId(v string) *AssumeUserInfo {
	s.Id = &v
	return s
}

func (s *AssumeUserInfo) SetSecurityToken(v string) *AssumeUserInfo {
	s.SecurityToken = &v
	return s
}

func (s *AssumeUserInfo) SetType(v string) *AssumeUserInfo {
	s.Type = &v
	return s
}

func (s *AssumeUserInfo) Validate() error {
	return dara.Validate(s)
}
