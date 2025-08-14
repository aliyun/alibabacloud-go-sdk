// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateEventRequest
	GetLang() *string
	SetCreateType(v string) *CreateEventRequest
	GetCreateType() *string
	SetEventName(v string) *CreateEventRequest
	GetEventName() *string
	SetInputFieldsStr(v string) *CreateEventRequest
	GetInputFieldsStr() *string
	SetMemo(v string) *CreateEventRequest
	GetMemo() *string
	SetRegId(v string) *CreateEventRequest
	GetRegId() *string
	SetTemplateCode(v string) *CreateEventRequest
	GetTemplateCode() *string
	SetTemplateName(v string) *CreateEventRequest
	GetTemplateName() *string
	SetTemplateType(v string) *CreateEventRequest
	GetTemplateType() *string
}

type CreateEventRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event name.
	//
	// example:
	//
	// 登录事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Input parameters, JSON string.
	//
	// example:
	//
	// [{"fieldCode":"accountId","description":"用户的账户ID，唯一标识一个账户的id","fieldRank":1,"title":"账户ID","fieldType":"STRING","fieldSource":"DEFAULT"},{"fieldCode":"hitRules","fieldRank":2,"title":"命中策略","fieldType":"STRING","fieldSource":"DEFAULT"},{"fieldCode":"age","description":"","fieldRank":3,"title":"年龄","fieldType":"INT","fieldSource":"DEFAULT"},{"fieldCode":"ip","description":"IP地址","fieldRank":4,"title":"IP地址","fieldType":"STRING","fieldSource":"DEFAULT"},{"fieldCode":"tags","fieldRank":5,"title":"风险标签","fieldType":"STRING","fieldSource":"DEFAULT"},{"fieldCode":"score","fieldRank":6,"title":"风险分值","fieldType":"DOUBLE","fieldSource":"DEFAULT"},{"fieldCode":"hitList","fieldRank":7}]
	InputFieldsStr *string `json:"inputFieldsStr,omitempty" xml:"inputFieldsStr,omitempty"`
	// Memo information
	//
	// example:
	//
	// 登录事件描述
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Input field template type
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Published template name.
	//
	// example:
	//
	// 注册事件模版
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// Template type.
	//
	// example:
	//
	// TASK
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s CreateEventRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRequest) GoString() string {
	return s.String()
}

func (s *CreateEventRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateEventRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *CreateEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *CreateEventRequest) GetInputFieldsStr() *string {
	return s.InputFieldsStr
}

func (s *CreateEventRequest) GetMemo() *string {
	return s.Memo
}

func (s *CreateEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateEventRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateEventRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateEventRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateEventRequest) SetLang(v string) *CreateEventRequest {
	s.Lang = &v
	return s
}

func (s *CreateEventRequest) SetCreateType(v string) *CreateEventRequest {
	s.CreateType = &v
	return s
}

func (s *CreateEventRequest) SetEventName(v string) *CreateEventRequest {
	s.EventName = &v
	return s
}

func (s *CreateEventRequest) SetInputFieldsStr(v string) *CreateEventRequest {
	s.InputFieldsStr = &v
	return s
}

func (s *CreateEventRequest) SetMemo(v string) *CreateEventRequest {
	s.Memo = &v
	return s
}

func (s *CreateEventRequest) SetRegId(v string) *CreateEventRequest {
	s.RegId = &v
	return s
}

func (s *CreateEventRequest) SetTemplateCode(v string) *CreateEventRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateEventRequest) SetTemplateName(v string) *CreateEventRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateEventRequest) SetTemplateType(v string) *CreateEventRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateEventRequest) Validate() error {
	return dara.Validate(s)
}
