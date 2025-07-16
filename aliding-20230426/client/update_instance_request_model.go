// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *UpdateInstanceRequest
	GetAppType() *string
	SetLanguage(v string) *UpdateInstanceRequest
	GetLanguage() *string
	SetProcessInstanceId(v string) *UpdateInstanceRequest
	GetProcessInstanceId() *string
	SetSystemToken(v string) *UpdateInstanceRequest
	GetSystemToken() *string
	SetUpdateFormDataJson(v string) *UpdateInstanceRequest
	GetUpdateFormDataJson() *string
}

type UpdateInstanceRequest struct {
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
	// f30233fb-72e1-4af4-8cb8-c7e0ea9ee530
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// hexxyy
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// example:
	//
	// {}
	UpdateFormDataJson *string `json:"UpdateFormDataJson,omitempty" xml:"UpdateFormDataJson,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAppType() *string {
	return s.AppType
}

func (s *UpdateInstanceRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateInstanceRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *UpdateInstanceRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *UpdateInstanceRequest) GetUpdateFormDataJson() *string {
	return s.UpdateFormDataJson
}

func (s *UpdateInstanceRequest) SetAppType(v string) *UpdateInstanceRequest {
	s.AppType = &v
	return s
}

func (s *UpdateInstanceRequest) SetLanguage(v string) *UpdateInstanceRequest {
	s.Language = &v
	return s
}

func (s *UpdateInstanceRequest) SetProcessInstanceId(v string) *UpdateInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetSystemToken(v string) *UpdateInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetUpdateFormDataJson(v string) *UpdateInstanceRequest {
	s.UpdateFormDataJson = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
