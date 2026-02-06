// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentUserSayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ListBeebotIntentUserSayRequest
	GetContent() *string
	SetInstanceId(v string) *ListBeebotIntentUserSayRequest
	GetInstanceId() *string
	SetIntentId(v string) *ListBeebotIntentUserSayRequest
	GetIntentId() *string
	SetPageNumber(v int32) *ListBeebotIntentUserSayRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBeebotIntentUserSayRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListBeebotIntentUserSayRequest
	GetScriptId() *string
}

type ListBeebotIntentUserSayRequest struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10717802
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListBeebotIntentUserSayRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentUserSayRequest) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentUserSayRequest) GetContent() *string {
	return s.Content
}

func (s *ListBeebotIntentUserSayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBeebotIntentUserSayRequest) GetIntentId() *string {
	return s.IntentId
}

func (s *ListBeebotIntentUserSayRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBeebotIntentUserSayRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBeebotIntentUserSayRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListBeebotIntentUserSayRequest) SetContent(v string) *ListBeebotIntentUserSayRequest {
	s.Content = &v
	return s
}

func (s *ListBeebotIntentUserSayRequest) SetInstanceId(v string) *ListBeebotIntentUserSayRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBeebotIntentUserSayRequest) SetIntentId(v string) *ListBeebotIntentUserSayRequest {
	s.IntentId = &v
	return s
}

func (s *ListBeebotIntentUserSayRequest) SetPageNumber(v int32) *ListBeebotIntentUserSayRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBeebotIntentUserSayRequest) SetPageSize(v int32) *ListBeebotIntentUserSayRequest {
	s.PageSize = &v
	return s
}

func (s *ListBeebotIntentUserSayRequest) SetScriptId(v string) *ListBeebotIntentUserSayRequest {
	s.ScriptId = &v
	return s
}

func (s *ListBeebotIntentUserSayRequest) Validate() error {
	return dara.Validate(s)
}
