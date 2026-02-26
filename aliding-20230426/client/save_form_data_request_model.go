// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *SaveFormDataRequest
	GetAppType() *string
	SetFormDataJson(v string) *SaveFormDataRequest
	GetFormDataJson() *string
	SetFormUuid(v string) *SaveFormDataRequest
	GetFormUuid() *string
	SetLanguage(v string) *SaveFormDataRequest
	GetLanguage() *string
	SetSystemToken(v string) *SaveFormDataRequest
	GetSystemToken() *string
}

type SaveFormDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {\\"textField_jcpm6agt\\": \\"单行\\",\\"employeeField_jcos0sar\\": [\\"workno\\"]}
	FormDataJson *string `json:"FormDataJson,omitempty" xml:"FormDataJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-EF6Y4G8WO2FN0SUB43TDQ3CGC3FMFQ1G9400RCJ3
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s SaveFormDataRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveFormDataRequest) GoString() string {
	return s.String()
}

func (s *SaveFormDataRequest) GetAppType() *string {
	return s.AppType
}

func (s *SaveFormDataRequest) GetFormDataJson() *string {
	return s.FormDataJson
}

func (s *SaveFormDataRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *SaveFormDataRequest) GetLanguage() *string {
	return s.Language
}

func (s *SaveFormDataRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *SaveFormDataRequest) SetAppType(v string) *SaveFormDataRequest {
	s.AppType = &v
	return s
}

func (s *SaveFormDataRequest) SetFormDataJson(v string) *SaveFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *SaveFormDataRequest) SetFormUuid(v string) *SaveFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *SaveFormDataRequest) SetLanguage(v string) *SaveFormDataRequest {
	s.Language = &v
	return s
}

func (s *SaveFormDataRequest) SetSystemToken(v string) *SaveFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *SaveFormDataRequest) Validate() error {
	return dara.Validate(s)
}
