// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *SdkGenerateByGroupRequest
	GetGroupId() *string
	SetLanguage(v string) *SdkGenerateByGroupRequest
	GetLanguage() *string
	SetSecurityToken(v string) *SdkGenerateByGroupRequest
	GetSecurityToken() *string
}

type SdkGenerateByGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1a991a450b9548a1a3df38fd3af117c2
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// java
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SdkGenerateByGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByGroupRequest) GoString() string {
	return s.String()
}

func (s *SdkGenerateByGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *SdkGenerateByGroupRequest) GetLanguage() *string {
	return s.Language
}

func (s *SdkGenerateByGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SdkGenerateByGroupRequest) SetGroupId(v string) *SdkGenerateByGroupRequest {
	s.GroupId = &v
	return s
}

func (s *SdkGenerateByGroupRequest) SetLanguage(v string) *SdkGenerateByGroupRequest {
	s.Language = &v
	return s
}

func (s *SdkGenerateByGroupRequest) SetSecurityToken(v string) *SdkGenerateByGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *SdkGenerateByGroupRequest) Validate() error {
	return dara.Validate(s)
}
