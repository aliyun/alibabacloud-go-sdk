// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *TerminateInstanceRequest
	GetAppType() *string
	SetLanguage(v string) *TerminateInstanceRequest
	GetLanguage() *string
	SetProcessInstanceId(v string) *TerminateInstanceRequest
	GetProcessInstanceId() *string
	SetSystemToken(v string) *TerminateInstanceRequest
	GetSystemToken() *string
}

type TerminateInstanceRequest struct {
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
}

func (s TerminateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TerminateInstanceRequest) GoString() string {
	return s.String()
}

func (s *TerminateInstanceRequest) GetAppType() *string {
	return s.AppType
}

func (s *TerminateInstanceRequest) GetLanguage() *string {
	return s.Language
}

func (s *TerminateInstanceRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *TerminateInstanceRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *TerminateInstanceRequest) SetAppType(v string) *TerminateInstanceRequest {
	s.AppType = &v
	return s
}

func (s *TerminateInstanceRequest) SetLanguage(v string) *TerminateInstanceRequest {
	s.Language = &v
	return s
}

func (s *TerminateInstanceRequest) SetProcessInstanceId(v string) *TerminateInstanceRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *TerminateInstanceRequest) SetSystemToken(v string) *TerminateInstanceRequest {
	s.SystemToken = &v
	return s
}

func (s *TerminateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
