// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptVoiceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptVoiceConfigsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListScriptVoiceConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListScriptVoiceConfigsRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListScriptVoiceConfigsRequest
	GetScriptId() *string
}

type ListScriptVoiceConfigsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f95c7ca6-872c-4765-8493-165a8dfc682d
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListScriptVoiceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVoiceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListScriptVoiceConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptVoiceConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptVoiceConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptVoiceConfigsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptVoiceConfigsRequest) SetInstanceId(v string) *ListScriptVoiceConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptVoiceConfigsRequest) SetPageNumber(v int32) *ListScriptVoiceConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScriptVoiceConfigsRequest) SetPageSize(v int32) *ListScriptVoiceConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListScriptVoiceConfigsRequest) SetScriptId(v string) *ListScriptVoiceConfigsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListScriptVoiceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
