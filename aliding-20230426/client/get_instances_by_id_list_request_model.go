// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancesByIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetInstancesByIdListRequest
	GetAppType() *string
	SetLanguage(v string) *GetInstancesByIdListRequest
	GetLanguage() *string
	SetProcessInstanceIds(v string) *GetInstancesByIdListRequest
	GetProcessInstanceIds() *string
	SetSystemToken(v string) *GetInstancesByIdListRequest
	GetSystemToken() *string
}

type GetInstancesByIdListRequest struct {
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
	// inst-123,inst-223
	ProcessInstanceIds *string `json:"ProcessInstanceIds,omitempty" xml:"ProcessInstanceIds,omitempty"`
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
}

func (s GetInstancesByIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstancesByIdListRequest) GoString() string {
	return s.String()
}

func (s *GetInstancesByIdListRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetInstancesByIdListRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetInstancesByIdListRequest) GetProcessInstanceIds() *string {
	return s.ProcessInstanceIds
}

func (s *GetInstancesByIdListRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetInstancesByIdListRequest) SetAppType(v string) *GetInstancesByIdListRequest {
	s.AppType = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetLanguage(v string) *GetInstancesByIdListRequest {
	s.Language = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetProcessInstanceIds(v string) *GetInstancesByIdListRequest {
	s.ProcessInstanceIds = &v
	return s
}

func (s *GetInstancesByIdListRequest) SetSystemToken(v string) *GetInstancesByIdListRequest {
	s.SystemToken = &v
	return s
}

func (s *GetInstancesByIdListRequest) Validate() error {
	return dara.Validate(s)
}
