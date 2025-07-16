// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateFormDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *CreateOrUpdateFormDataRequest
	GetAppType() *string
	SetFormDataJson(v string) *CreateOrUpdateFormDataRequest
	GetFormDataJson() *string
	SetFormUuid(v string) *CreateOrUpdateFormDataRequest
	GetFormUuid() *string
	SetNoExecuteExpression(v bool) *CreateOrUpdateFormDataRequest
	GetNoExecuteExpression() *bool
	SetSearchCondition(v string) *CreateOrUpdateFormDataRequest
	GetSearchCondition() *string
	SetSystemToken(v string) *CreateOrUpdateFormDataRequest
	GetSystemToken() *string
	SetUserId(v string) *CreateOrUpdateFormDataRequest
	GetUserId() *string
}

type CreateOrUpdateFormDataRequest struct {
	// example:
	//
	// APP_XCE0EVXS6DYG3YDYC5RD
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// {\\"countrySelectField_l0c1cwiu\\":[{\\"value\\":\\"US\\"}]}
	FormDataJson *string `json:"FormDataJson,omitempty" xml:"FormDataJson,omitempty"`
	// example:
	//
	// FORM-GX866MC1NC1VOFF6WVQW33FD16E23L3CPMKVKA
	FormUuid *string `json:"FormUuid,omitempty" xml:"FormUuid,omitempty"`
	// example:
	//
	// false
	NoExecuteExpression *bool   `json:"NoExecuteExpression,omitempty" xml:"NoExecuteExpression,omitempty"`
	SearchCondition     *string `json:"SearchCondition,omitempty" xml:"SearchCondition,omitempty"`
	// example:
	//
	// 09866181UTZVVD4R3DC955FNKIM52HVPU5WWK7
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateOrUpdateFormDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateFormDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataRequest) GetAppType() *string {
	return s.AppType
}

func (s *CreateOrUpdateFormDataRequest) GetFormDataJson() *string {
	return s.FormDataJson
}

func (s *CreateOrUpdateFormDataRequest) GetFormUuid() *string {
	return s.FormUuid
}

func (s *CreateOrUpdateFormDataRequest) GetNoExecuteExpression() *bool {
	return s.NoExecuteExpression
}

func (s *CreateOrUpdateFormDataRequest) GetSearchCondition() *string {
	return s.SearchCondition
}

func (s *CreateOrUpdateFormDataRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *CreateOrUpdateFormDataRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateOrUpdateFormDataRequest) SetAppType(v string) *CreateOrUpdateFormDataRequest {
	s.AppType = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormDataJson(v string) *CreateOrUpdateFormDataRequest {
	s.FormDataJson = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetFormUuid(v string) *CreateOrUpdateFormDataRequest {
	s.FormUuid = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetNoExecuteExpression(v bool) *CreateOrUpdateFormDataRequest {
	s.NoExecuteExpression = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSearchCondition(v string) *CreateOrUpdateFormDataRequest {
	s.SearchCondition = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetSystemToken(v string) *CreateOrUpdateFormDataRequest {
	s.SystemToken = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) SetUserId(v string) *CreateOrUpdateFormDataRequest {
	s.UserId = &v
	return s
}

func (s *CreateOrUpdateFormDataRequest) Validate() error {
	return dara.Validate(s)
}
