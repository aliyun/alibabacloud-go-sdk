// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackPathEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v int64) *GetAttackPathEventDetailRequest
	GetEventId() *int64
	SetEventSource(v string) *GetAttackPathEventDetailRequest
	GetEventSource() *string
	SetLang(v string) *GetAttackPathEventDetailRequest
	GetLang() *string
}

type GetAttackPathEventDetailRequest struct {
	// Event ID.
	//
	// > You can call [ListAttackPathEvent](~~ListAttackPathEvent~~) to query the event ID.
	//
	// example:
	//
	// 123
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// 数据来源。默认值为**default**。取值：
	//
	// - **caasm**：攻击面
	//
	// - **default**：攻击路径
	//
	// example:
	//
	// default
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The language type for request and response, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetAttackPathEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackPathEventDetailRequest) GoString() string {
	return s.String()
}

func (s *GetAttackPathEventDetailRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *GetAttackPathEventDetailRequest) GetEventSource() *string {
	return s.EventSource
}

func (s *GetAttackPathEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *GetAttackPathEventDetailRequest) SetEventId(v int64) *GetAttackPathEventDetailRequest {
	s.EventId = &v
	return s
}

func (s *GetAttackPathEventDetailRequest) SetEventSource(v string) *GetAttackPathEventDetailRequest {
	s.EventSource = &v
	return s
}

func (s *GetAttackPathEventDetailRequest) SetLang(v string) *GetAttackPathEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *GetAttackPathEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
