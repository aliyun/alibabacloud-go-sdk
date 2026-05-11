// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAsrConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *ModifyAsrConfigRequest
	GetAppKey() *string
	SetAsrAcousticModelId(v string) *ModifyAsrConfigRequest
	GetAsrAcousticModelId() *string
	SetAsrClassVocabularyId(v string) *ModifyAsrConfigRequest
	GetAsrClassVocabularyId() *string
	SetAsrCustomizationId(v string) *ModifyAsrConfigRequest
	GetAsrCustomizationId() *string
	SetAsrOverrides(v string) *ModifyAsrConfigRequest
	GetAsrOverrides() *string
	SetAsrVocabularyId(v string) *ModifyAsrConfigRequest
	GetAsrVocabularyId() *string
	SetConfigLevel(v int32) *ModifyAsrConfigRequest
	GetConfigLevel() *int32
	SetEngine(v string) *ModifyAsrConfigRequest
	GetEngine() *string
	SetEntryId(v string) *ModifyAsrConfigRequest
	GetEntryId() *string
	SetNlsServiceType(v string) *ModifyAsrConfigRequest
	GetNlsServiceType() *string
}

type ModifyAsrConfigRequest struct {
	// example:
	//
	// your_app_key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrAcousticModelId *string `json:"AsrAcousticModelId,omitempty" xml:"AsrAcousticModelId,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrClassVocabularyId *string `json:"AsrClassVocabularyId,omitempty" xml:"AsrClassVocabularyId,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrCustomizationId *string `json:"AsrCustomizationId,omitempty" xml:"AsrCustomizationId,omitempty"`
	AsrOverrides       *string `json:"AsrOverrides,omitempty" xml:"AsrOverrides,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// example:
	//
	// 0
	ConfigLevel *int32  `json:"ConfigLevel,omitempty" xml:"ConfigLevel,omitempty"`
	Engine      *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 6cc9f5ca-2cb6-4cc7-a46b-2bbfd3e61b22
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// example:
	//
	// 语音识别
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
}

func (s ModifyAsrConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAsrConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAsrConfigRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ModifyAsrConfigRequest) GetAsrAcousticModelId() *string {
	return s.AsrAcousticModelId
}

func (s *ModifyAsrConfigRequest) GetAsrClassVocabularyId() *string {
	return s.AsrClassVocabularyId
}

func (s *ModifyAsrConfigRequest) GetAsrCustomizationId() *string {
	return s.AsrCustomizationId
}

func (s *ModifyAsrConfigRequest) GetAsrOverrides() *string {
	return s.AsrOverrides
}

func (s *ModifyAsrConfigRequest) GetAsrVocabularyId() *string {
	return s.AsrVocabularyId
}

func (s *ModifyAsrConfigRequest) GetConfigLevel() *int32 {
	return s.ConfigLevel
}

func (s *ModifyAsrConfigRequest) GetEngine() *string {
	return s.Engine
}

func (s *ModifyAsrConfigRequest) GetEntryId() *string {
	return s.EntryId
}

func (s *ModifyAsrConfigRequest) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *ModifyAsrConfigRequest) SetAppKey(v string) *ModifyAsrConfigRequest {
	s.AppKey = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrAcousticModelId(v string) *ModifyAsrConfigRequest {
	s.AsrAcousticModelId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrClassVocabularyId(v string) *ModifyAsrConfigRequest {
	s.AsrClassVocabularyId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrCustomizationId(v string) *ModifyAsrConfigRequest {
	s.AsrCustomizationId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrOverrides(v string) *ModifyAsrConfigRequest {
	s.AsrOverrides = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetAsrVocabularyId(v string) *ModifyAsrConfigRequest {
	s.AsrVocabularyId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetConfigLevel(v int32) *ModifyAsrConfigRequest {
	s.ConfigLevel = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetEngine(v string) *ModifyAsrConfigRequest {
	s.Engine = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetEntryId(v string) *ModifyAsrConfigRequest {
	s.EntryId = &v
	return s
}

func (s *ModifyAsrConfigRequest) SetNlsServiceType(v string) *ModifyAsrConfigRequest {
	s.NlsServiceType = &v
	return s
}

func (s *ModifyAsrConfigRequest) Validate() error {
	return dara.Validate(s)
}
