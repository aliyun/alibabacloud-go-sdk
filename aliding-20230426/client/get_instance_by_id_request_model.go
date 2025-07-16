// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetInstanceByIdRequest
	GetAppType() *string
	SetId(v string) *GetInstanceByIdRequest
	GetId() *string
	SetLanguage(v string) *GetInstanceByIdRequest
	GetLanguage() *string
	SetSystemToken(v string) *GetInstanceByIdRequest
	GetSystemToken() *string
}

type GetInstanceByIdRequest struct {
	// example:
	//
	// APP_PBxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// f30233fb-72e1-xxx
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

func (s GetInstanceByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetInstanceByIdRequest) GetId() *string {
	return s.Id
}

func (s *GetInstanceByIdRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetInstanceByIdRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetInstanceByIdRequest) SetAppType(v string) *GetInstanceByIdRequest {
	s.AppType = &v
	return s
}

func (s *GetInstanceByIdRequest) SetId(v string) *GetInstanceByIdRequest {
	s.Id = &v
	return s
}

func (s *GetInstanceByIdRequest) SetLanguage(v string) *GetInstanceByIdRequest {
	s.Language = &v
	return s
}

func (s *GetInstanceByIdRequest) SetSystemToken(v string) *GetInstanceByIdRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstanceByIdRequest) Validate() error {
	return dara.Validate(s)
}
