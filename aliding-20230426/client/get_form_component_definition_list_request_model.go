// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormComponentDefinitionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetFormComponentDefinitionListRequest
	GetAppType() *string
	SetFormUuid(v string) *GetFormComponentDefinitionListRequest
	GetFormUuid() *string
	SetLanguage(v string) *GetFormComponentDefinitionListRequest
	GetLanguage() *string
	SetSystemToken(v string) *GetFormComponentDefinitionListRequest
	GetSystemToken() *string
}

type GetFormComponentDefinitionListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKT0xxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FORM-xxxxx
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

func (s GetFormComponentDefinitionListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListRequest) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetFormComponentDefinitionListRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *GetFormComponentDefinitionListRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetFormComponentDefinitionListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetFormComponentDefinitionListRequest) SetAppType(v string) *GetFormComponentDefinitionListRequest {
	s.AppType = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) SetFormUuid(v string) *GetFormComponentDefinitionListRequest {
	s.FormUuid = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) SetLanguage(v string) *GetFormComponentDefinitionListRequest {
	s.Language = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) SetSystemToken(v string) *GetFormComponentDefinitionListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetFormComponentDefinitionListRequest) Validate() error {
	return dara.Validate(s)
}
