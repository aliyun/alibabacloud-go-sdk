// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *DeleteFormDataRequest
	GetAppType() *string
	SetFormInstanceId(v string) *DeleteFormDataRequest
	GetFormInstanceId() *string
	SetLanguage(v string) *DeleteFormDataRequest
	GetLanguage() *string
	SetSystemToken(v string) *DeleteFormDataRequest
	GetSystemToken() *string
}

type DeleteFormDataRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKTxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM_INST_12345
	FormInstanceId *string `json:"FormInstanceId,omitempty" xml:"FormInstanceId,omitempty"`
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

func (s DeleteFormDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteFormDataRequest) GetAppType() *string {
	return s.AppType
}

func (s *DeleteFormDataRequest) GetFormInstanceId() *string {
	return s.FormInstanceId
}

func (s *DeleteFormDataRequest) GetLanguage() *string {
	return s.Language
}

func (s *DeleteFormDataRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *DeleteFormDataRequest) SetAppType(v string) *DeleteFormDataRequest {
	s.AppType = &v
	return s
}

func (s *DeleteFormDataRequest) SetFormInstanceId(v string) *DeleteFormDataRequest {
	s.FormInstanceId = &v
	return s
}

func (s *DeleteFormDataRequest) SetLanguage(v string) *DeleteFormDataRequest {
	s.Language = &v
	return s
}

func (s *DeleteFormDataRequest) SetSystemToken(v string) *DeleteFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *DeleteFormDataRequest) Validate() error {
	return dara.Validate(s)
}
