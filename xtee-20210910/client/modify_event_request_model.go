// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyEventRequest
	GetLang() *string
	SetBizVersion(v int32) *ModifyEventRequest
	GetBizVersion() *int32
	SetCreateType(v string) *ModifyEventRequest
	GetCreateType() *string
	SetEventCode(v string) *ModifyEventRequest
	GetEventCode() *string
	SetEventName(v string) *ModifyEventRequest
	GetEventName() *string
	SetInputFieldsStr(v string) *ModifyEventRequest
	GetInputFieldsStr() *string
	SetMemo(v string) *ModifyEventRequest
	GetMemo() *string
	SetRegId(v string) *ModifyEventRequest
	GetRegId() *string
	SetTemplateType(v string) *ModifyEventRequest
	GetTemplateType() *string
}

type ModifyEventRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Version number (latest).
	//
	// example:
	//
	// 1
	BizVersion *int32 `json:"bizVersion,omitempty" xml:"bizVersion,omitempty"`
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_ambiby3420
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Input parameters, JSON string.
	//
	// example:
	//
	// [{\\"fieldCode\\":\\"sessionId\\",\\"description\\":\\"会话ID\\",\\"fieldRank\\":0,\\"title\\":\\"会话ID\\",\\"fieldType\\":\\"STRING\\",\\"fieldSource\\":\\"DEFAULT\\"},{\\"fieldCode\\":\\"tags\\",\\"fieldRank\\":1,\\"title\\":\\"风险标签\\",\\"fieldType\\":\\"STRING\\",\\"fieldSource\\":\\"DEFAULT\\"},{\\"fieldCode\\":\\"score\\",\\"fieldRank\\":2,\\"title\\":\\"风险分值\\",\\"fieldType\\":\\"DOUBLE\\",\\"fieldSource\\":\\"DEFAULT\\"},{\\"fieldCode\\":\\"hitRules\\",\\"fieldRank\\":3,\\"title\\":\\"命中策略\\",\\"fieldType\\":\\"STRING\\",\\"fieldSource\\":\\"DEFAULT\\"}]
	InputFieldsStr *string `json:"inputFieldsStr,omitempty" xml:"inputFieldsStr,omitempty"`
	// Memo.
	//
	// example:
	//
	// 备注
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Template type
	//
	// example:
	//
	// 暂无
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s ModifyEventRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyEventRequest) GetBizVersion() *int32 {
	return s.BizVersion
}

func (s *ModifyEventRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *ModifyEventRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *ModifyEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *ModifyEventRequest) GetInputFieldsStr() *string {
	return s.InputFieldsStr
}

func (s *ModifyEventRequest) GetMemo() *string {
	return s.Memo
}

func (s *ModifyEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyEventRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ModifyEventRequest) SetLang(v string) *ModifyEventRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventRequest) SetBizVersion(v int32) *ModifyEventRequest {
	s.BizVersion = &v
	return s
}

func (s *ModifyEventRequest) SetCreateType(v string) *ModifyEventRequest {
	s.CreateType = &v
	return s
}

func (s *ModifyEventRequest) SetEventCode(v string) *ModifyEventRequest {
	s.EventCode = &v
	return s
}

func (s *ModifyEventRequest) SetEventName(v string) *ModifyEventRequest {
	s.EventName = &v
	return s
}

func (s *ModifyEventRequest) SetInputFieldsStr(v string) *ModifyEventRequest {
	s.InputFieldsStr = &v
	return s
}

func (s *ModifyEventRequest) SetMemo(v string) *ModifyEventRequest {
	s.Memo = &v
	return s
}

func (s *ModifyEventRequest) SetRegId(v string) *ModifyEventRequest {
	s.RegId = &v
	return s
}

func (s *ModifyEventRequest) SetTemplateType(v string) *ModifyEventRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifyEventRequest) Validate() error {
	return dara.Validate(s)
}
