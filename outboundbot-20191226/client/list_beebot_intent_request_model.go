// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBeebotIntentRequest
	GetInstanceId() *string
	SetIntentName(v string) *ListBeebotIntentRequest
	GetIntentName() *string
	SetPageNumber(v int32) *ListBeebotIntentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBeebotIntentRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListBeebotIntentRequest
	GetScriptId() *string
}

type ListBeebotIntentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
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

func (s ListBeebotIntentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentRequest) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBeebotIntentRequest) GetIntentName() *string {
	return s.IntentName
}

func (s *ListBeebotIntentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBeebotIntentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBeebotIntentRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListBeebotIntentRequest) SetInstanceId(v string) *ListBeebotIntentRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBeebotIntentRequest) SetIntentName(v string) *ListBeebotIntentRequest {
	s.IntentName = &v
	return s
}

func (s *ListBeebotIntentRequest) SetPageNumber(v int32) *ListBeebotIntentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBeebotIntentRequest) SetPageSize(v int32) *ListBeebotIntentRequest {
	s.PageSize = &v
	return s
}

func (s *ListBeebotIntentRequest) SetScriptId(v string) *ListBeebotIntentRequest {
	s.ScriptId = &v
	return s
}

func (s *ListBeebotIntentRequest) Validate() error {
	return dara.Validate(s)
}
