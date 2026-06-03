// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentLgfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBeebotIntentLgfRequest
	GetInstanceId() *string
	SetIntentId(v string) *ListBeebotIntentLgfRequest
	GetIntentId() *string
	SetLgfText(v string) *ListBeebotIntentLgfRequest
	GetLgfText() *string
	SetPageNumber(v int32) *ListBeebotIntentLgfRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBeebotIntentLgfRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListBeebotIntentLgfRequest
	GetScriptId() *string
}

type ListBeebotIntentLgfRequest struct {
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
	LgfText  *string `json:"LgfText,omitempty" xml:"LgfText,omitempty"`
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

func (s ListBeebotIntentLgfRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentLgfRequest) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentLgfRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBeebotIntentLgfRequest) GetIntentId() *string {
	return s.IntentId
}

func (s *ListBeebotIntentLgfRequest) GetLgfText() *string {
	return s.LgfText
}

func (s *ListBeebotIntentLgfRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBeebotIntentLgfRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBeebotIntentLgfRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListBeebotIntentLgfRequest) SetInstanceId(v string) *ListBeebotIntentLgfRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBeebotIntentLgfRequest) SetIntentId(v string) *ListBeebotIntentLgfRequest {
	s.IntentId = &v
	return s
}

func (s *ListBeebotIntentLgfRequest) SetLgfText(v string) *ListBeebotIntentLgfRequest {
	s.LgfText = &v
	return s
}

func (s *ListBeebotIntentLgfRequest) SetPageNumber(v int32) *ListBeebotIntentLgfRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBeebotIntentLgfRequest) SetPageSize(v int32) *ListBeebotIntentLgfRequest {
	s.PageSize = &v
	return s
}

func (s *ListBeebotIntentLgfRequest) SetScriptId(v string) *ListBeebotIntentLgfRequest {
	s.ScriptId = &v
	return s
}

func (s *ListBeebotIntentLgfRequest) Validate() error {
	return dara.Validate(s)
}
