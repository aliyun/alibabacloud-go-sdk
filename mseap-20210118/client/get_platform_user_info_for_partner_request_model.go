// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlatformUserInfoForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetPlatformUserInfoForPartnerRequest
	GetAppId() *string
	SetPlatformType(v string) *GetPlatformUserInfoForPartnerRequest
	GetPlatformType() *string
	SetUserId(v string) *GetPlatformUserInfoForPartnerRequest
	GetUserId() *string
}

type GetPlatformUserInfoForPartnerRequest struct {
	// example:
	//
	// app-0wceagu85ceaaais
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// MP
	PlatformType *string `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// example:
	//
	// 123153124411
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetPlatformUserInfoForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlatformUserInfoForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetPlatformUserInfoForPartnerRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetPlatformUserInfoForPartnerRequest) GetPlatformType() *string {
	return s.PlatformType
}

func (s *GetPlatformUserInfoForPartnerRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetPlatformUserInfoForPartnerRequest) SetAppId(v string) *GetPlatformUserInfoForPartnerRequest {
	s.AppId = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerRequest) SetPlatformType(v string) *GetPlatformUserInfoForPartnerRequest {
	s.PlatformType = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerRequest) SetUserId(v string) *GetPlatformUserInfoForPartnerRequest {
	s.UserId = &v
	return s
}

func (s *GetPlatformUserInfoForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
