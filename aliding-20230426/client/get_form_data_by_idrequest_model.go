// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormDataByIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetFormDataByIDRequest
	GetAppType() *string
	SetId(v string) *GetFormDataByIDRequest
	GetId() *string
	SetLanguage(v string) *GetFormDataByIDRequest
	GetLanguage() *string
	SetSystemToken(v string) *GetFormDataByIDRequest
	GetSystemToken() *string
}

type GetFormDataByIDRequest struct {
	// example:
	//
	// APP_PBKT0MFBEBTDO8T7SLVP
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// FORM_INST_12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetFormDataByIDRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDRequest) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetFormDataByIDRequest) GetId() *string {
	return s.Id
}

func (s *GetFormDataByIDRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetFormDataByIDRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetFormDataByIDRequest) SetAppType(v string) *GetFormDataByIDRequest {
	s.AppType = &v
	return s
}

func (s *GetFormDataByIDRequest) SetId(v string) *GetFormDataByIDRequest {
	s.Id = &v
	return s
}

func (s *GetFormDataByIDRequest) SetLanguage(v string) *GetFormDataByIDRequest {
	s.Language = &v
	return s
}

func (s *GetFormDataByIDRequest) SetSystemToken(v string) *GetFormDataByIDRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormDataByIDRequest) Validate() error {
	return dara.Validate(s)
}
