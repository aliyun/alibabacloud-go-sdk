// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptVoiceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScriptVoiceConfigsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScriptVoiceConfigsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListScriptVoiceConfigsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScriptVoiceConfigsResponseBody
	GetRequestId() *string
	SetScriptVoiceConfigs(v *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) *ListScriptVoiceConfigsResponseBody
	GetScriptVoiceConfigs() *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs
	SetSuccess(v bool) *ListScriptVoiceConfigsResponseBody
	GetSuccess() *bool
}

type ListScriptVoiceConfigsResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of script voice recordings
	ScriptVoiceConfigs *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs `json:"ScriptVoiceConfigs,omitempty" xml:"ScriptVoiceConfigs,omitempty" type:"Struct"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScriptVoiceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVoiceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScriptVoiceConfigsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScriptVoiceConfigsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScriptVoiceConfigsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScriptVoiceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScriptVoiceConfigsResponseBody) GetScriptVoiceConfigs() *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs {
	return s.ScriptVoiceConfigs
}

func (s *ListScriptVoiceConfigsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScriptVoiceConfigsResponseBody) SetCode(v string) *ListScriptVoiceConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBody) SetHttpStatusCode(v int32) *ListScriptVoiceConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBody) SetMessage(v string) *ListScriptVoiceConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBody) SetRequestId(v string) *ListScriptVoiceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBody) SetScriptVoiceConfigs(v *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) *ListScriptVoiceConfigsResponseBody {
	s.ScriptVoiceConfigs = v
	return s
}

func (s *ListScriptVoiceConfigsResponseBody) SetSuccess(v bool) *ListScriptVoiceConfigsResponseBody {
	s.Success = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBody) Validate() error {
	if s.ScriptVoiceConfigs != nil {
		if err := s.ScriptVoiceConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs struct {
	// Array of script voice recordings
	List []*ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) GoString() string {
	return s.String()
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) GetList() []*ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	return s.List
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) SetList(v []*ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs {
	s.List = v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) SetPageNumber(v int32) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs {
	s.PageNumber = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) SetPageSize(v int32) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs {
	s.PageSize = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) SetTotalCount(v int32) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs {
	s.TotalCount = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigs) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList struct {
	// instance ID
	//
	// example:
	//
	// bdd49242-114c-4045-b1d1-25ccc1756c75
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script text content
	//
	// example:
	//
	// 请问你是 @name 吗
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// script ID
	//
	// example:
	//
	// a7441a05-43bb-4a2d-acb0-365f245d7a5b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// voice configuration ID
	//
	// example:
	//
	// 2c8fa91f-9856-4145-90f2-08252f09bc18
	ScriptVoiceConfigId *string `json:"ScriptVoiceConfigId,omitempty" xml:"ScriptVoiceConfigId,omitempty"`
	// relational data of script audio configuration
	//
	// example:
	//
	// ""
	ScriptWaveformRelation *string `json:"ScriptWaveformRelation,omitempty" xml:"ScriptWaveformRelation,omitempty"`
	// script source
	//
	// example:
	//
	// DIALOGUE_FLOW
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Voice type
	//
	// example:
	//
	// TTS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GoString() string {
	return s.String()
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetScriptVoiceConfigId() *string {
	return s.ScriptVoiceConfigId
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetScriptWaveformRelation() *string {
	return s.ScriptWaveformRelation
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetSource() *string {
	return s.Source
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) GetType() *string {
	return s.Type
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetInstanceId(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.InstanceId = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetScriptContent(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.ScriptContent = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetScriptId(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.ScriptId = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetScriptVoiceConfigId(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.ScriptVoiceConfigId = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetScriptWaveformRelation(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.ScriptWaveformRelation = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetSource(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.Source = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) SetType(v string) *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList {
	s.Type = &v
	return s
}

func (s *ListScriptVoiceConfigsResponseBodyScriptVoiceConfigsList) Validate() error {
	return dara.Validate(s)
}
