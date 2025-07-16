// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetProcessDefinitionRequest
	GetAppType() *string
	SetCorpId(v string) *GetProcessDefinitionRequest
	GetCorpId() *string
	SetGroupId(v string) *GetProcessDefinitionRequest
	GetGroupId() *string
	SetLanguage(v string) *GetProcessDefinitionRequest
	GetLanguage() *string
	SetNameSpace(v string) *GetProcessDefinitionRequest
	GetNameSpace() *string
	SetOrderNumber(v string) *GetProcessDefinitionRequest
	GetOrderNumber() *string
	SetProcessInstanceId(v string) *GetProcessDefinitionRequest
	GetProcessInstanceId() *string
	SetSystemToken(v string) *GetProcessDefinitionRequest
	GetSystemToken() *string
	SetSystemType(v string) *GetProcessDefinitionRequest
	GetSystemType() *string
}

type GetProcessDefinitionRequest struct {
	// example:
	//
	// APP_PBxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// ding123
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// c-xxafafaf
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// zh_CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// dingtalk
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// example:
	//
	// o-YDJKIxxx
	OrderNumber *string `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
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
	// hexabc1234
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s GetProcessDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetProcessDefinitionRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GetProcessDefinitionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *GetProcessDefinitionRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetProcessDefinitionRequest) GetNameSpace() *string {
	return s.NameSpace
}

func (s *GetProcessDefinitionRequest) GetOrderNumber() *string {
	return s.OrderNumber
}

func (s *GetProcessDefinitionRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *GetProcessDefinitionRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *GetProcessDefinitionRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *GetProcessDefinitionRequest) SetAppType(v string) *GetProcessDefinitionRequest {
	s.AppType = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetCorpId(v string) *GetProcessDefinitionRequest {
	s.CorpId = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetGroupId(v string) *GetProcessDefinitionRequest {
	s.GroupId = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetLanguage(v string) *GetProcessDefinitionRequest {
	s.Language = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetNameSpace(v string) *GetProcessDefinitionRequest {
	s.NameSpace = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetOrderNumber(v string) *GetProcessDefinitionRequest {
	s.OrderNumber = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetProcessInstanceId(v string) *GetProcessDefinitionRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetSystemToken(v string) *GetProcessDefinitionRequest {
	s.SystemToken = &v
	return s
}

func (s *GetProcessDefinitionRequest) SetSystemType(v string) *GetProcessDefinitionRequest {
	s.SystemType = &v
	return s
}

func (s *GetProcessDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
