// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdkGenerateByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *SdkGenerateByAppRequest
	GetAppId() *int64
	SetLanguage(v string) *SdkGenerateByAppRequest
	GetLanguage() *string
	SetSecurityToken(v string) *SdkGenerateByAppRequest
	GetSecurityToken() *string
}

type SdkGenerateByAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 110797019
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// java
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SdkGenerateByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s SdkGenerateByAppRequest) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *SdkGenerateByAppRequest) GetLanguage() *string {
	return s.Language
}

func (s *SdkGenerateByAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SdkGenerateByAppRequest) SetAppId(v int64) *SdkGenerateByAppRequest {
	s.AppId = &v
	return s
}

func (s *SdkGenerateByAppRequest) SetLanguage(v string) *SdkGenerateByAppRequest {
	s.Language = &v
	return s
}

func (s *SdkGenerateByAppRequest) SetSecurityToken(v string) *SdkGenerateByAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *SdkGenerateByAppRequest) Validate() error {
	return dara.Validate(s)
}
