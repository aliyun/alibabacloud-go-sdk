// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNavigationByFormTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *ListNavigationByFormTypeRequest
	GetAppType() *string
	SetFormType(v string) *ListNavigationByFormTypeRequest
	GetFormType() *string
	SetLanguage(v string) *ListNavigationByFormTypeRequest
	GetLanguage() *string
	SetSystemToken(v string) *ListNavigationByFormTypeRequest
	GetSystemToken() *string
}

type ListNavigationByFormTypeRequest struct {
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
	// receipt
	FormType *string `json:"FormType,omitempty" xml:"FormType,omitempty"`
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

func (s ListNavigationByFormTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNavigationByFormTypeRequest) GoString() string {
	return s.String()
}

func (s *ListNavigationByFormTypeRequest) GetAppType() *string {
	return s.AppType
}

func (s *ListNavigationByFormTypeRequest) GetFormType() *string {
	return s.FormType
}

func (s *ListNavigationByFormTypeRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListNavigationByFormTypeRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *ListNavigationByFormTypeRequest) SetAppType(v string) *ListNavigationByFormTypeRequest {
	s.AppType = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetFormType(v string) *ListNavigationByFormTypeRequest {
	s.FormType = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetLanguage(v string) *ListNavigationByFormTypeRequest {
	s.Language = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) SetSystemToken(v string) *ListNavigationByFormTypeRequest {
	s.SystemToken = &v
	return s
}

func (s *ListNavigationByFormTypeRequest) Validate() error {
	return dara.Validate(s)
}
