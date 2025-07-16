// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActivityListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetActivityListRequest
	GetAppType() *string
	SetLanguage(v string) *GetActivityListRequest
	GetLanguage() *string
	SetProcessCode(v string) *GetActivityListRequest
	GetProcessCode() *string
	SetSystemToken(v string) *GetActivityListRequest
	GetSystemToken() *string
}

type GetActivityListRequest struct {
	// example:
	//
	// APP_PBxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// TPROC--X1Gxx
	ProcessCode *string `json:"ProcessCode,omitempty" xml:"ProcessCode,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetActivityListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetActivityListRequest) GoString() string {
	return s.String()
}

func (s *GetActivityListRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetActivityListRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetActivityListRequest) GetProcessCode() *string {
	return s.ProcessCode
}

func (s *GetActivityListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetActivityListRequest) SetAppType(v string) *GetActivityListRequest {
	s.AppType = &v
	return s
}

func (s *GetActivityListRequest) SetLanguage(v string) *GetActivityListRequest {
	s.Language = &v
	return s
}

func (s *GetActivityListRequest) SetProcessCode(v string) *GetActivityListRequest {
	s.ProcessCode = &v
	return s
}

func (s *GetActivityListRequest) SetSystemToken(v string) *GetActivityListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetActivityListRequest) Validate() error {
	return dara.Validate(s)
}
