// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *UpdateFormDataRequest
	GetAppType() *string
	SetFormInstanceId(v string) *UpdateFormDataRequest
	GetFormInstanceId() *string
	SetLanguage(v string) *UpdateFormDataRequest
	GetLanguage() *string
	SetSystemToken(v string) *UpdateFormDataRequest
	GetSystemToken() *string
	SetUpdateFormDataJson(v string) *UpdateFormDataRequest
	GetUpdateFormDataJson() *string
	SetUseLatestVersion(v bool) *UpdateFormDataRequest
	GetUseLatestVersion() *bool
}

type UpdateFormDataRequest struct {
	// example:
	//
	// APP_PBKTxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// FORM_INxxx
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// example:
	//
	// {}
	UpdateFormDataJson *string `json:"UpdateFormDataJson,omitempty" xml:"UpdateFormDataJson,omitempty"`
	// example:
	//
	// false
	UseLatestVersion *bool `json:"UseLatestVersion,omitempty" xml:"UseLatestVersion,omitempty"`
}

func (s UpdateFormDataRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *UpdateFormDataRequest) GetAppType() *string {
	return s.AppType
}

func (s *UpdateFormDataRequest) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *UpdateFormDataRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateFormDataRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *UpdateFormDataRequest) GetUpdateFormDataJson() *string {
	return s.UpdateFormDataJson
}

func (s *UpdateFormDataRequest) GetUseLatestVersion() *bool {
	return s.UseLatestVersion
}

func (s *UpdateFormDataRequest) SetAppType(v string) *UpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *UpdateFormDataRequest) SetFormInstanceId(v string) *UpdateFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *UpdateFormDataRequest) SetLanguage(v string) *UpdateFormDataRequest {
	s.Language = &v
	return s
}

func (s *UpdateFormDataRequest) SetSystemToken(v string) *UpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateFormDataRequest) SetUpdateFormDataJson(v string) *UpdateFormDataRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateFormDataRequest) SetUseLatestVersion(v bool) *UpdateFormDataRequest {
	s.UseLatestVersion = &v
	return s
}

func (s *UpdateFormDataRequest) Validate() error {
	return dara.Validate(s)
}
