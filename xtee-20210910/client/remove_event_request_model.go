// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *RemoveEventRequest
	GetLang() *string
	SetCreateType(v string) *RemoveEventRequest
	GetCreateType() *string
	SetEventCode(v string) *RemoveEventRequest
	GetEventCode() *string
	SetId(v int64) *RemoveEventRequest
	GetId() *int64
	SetRegId(v string) *RemoveEventRequest
	GetRegId() *string
	SetTemplateCode(v string) *RemoveEventRequest
	GetTemplateCode() *string
}

type RemoveEventRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// The event code.
	//
	// example:
	//
	// de_arqbuy7206
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 2556
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The templatetype of the input parameter.
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s RemoveEventRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveEventRequest) GoString() string {
	return s.String()
}

func (s *RemoveEventRequest) GetLang() *string {
	return s.Lang
}

func (s *RemoveEventRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *RemoveEventRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *RemoveEventRequest) GetId() *int64 {
	return s.Id
}

func (s *RemoveEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *RemoveEventRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *RemoveEventRequest) SetLang(v string) *RemoveEventRequest {
	s.Lang = &v
	return s
}

func (s *RemoveEventRequest) SetCreateType(v string) *RemoveEventRequest {
	s.CreateType = &v
	return s
}

func (s *RemoveEventRequest) SetEventCode(v string) *RemoveEventRequest {
	s.EventCode = &v
	return s
}

func (s *RemoveEventRequest) SetId(v int64) *RemoveEventRequest {
	s.Id = &v
	return s
}

func (s *RemoveEventRequest) SetRegId(v string) *RemoveEventRequest {
	s.RegId = &v
	return s
}

func (s *RemoveEventRequest) SetTemplateCode(v string) *RemoveEventRequest {
	s.TemplateCode = &v
	return s
}

func (s *RemoveEventRequest) Validate() error {
	return dara.Validate(s)
}
