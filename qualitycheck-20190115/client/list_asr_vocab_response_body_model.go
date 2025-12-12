// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsrVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAsrVocabResponseBody
	GetCode() *string
	SetData(v *ListAsrVocabResponseBodyData) *ListAsrVocabResponseBody
	GetData() *ListAsrVocabResponseBodyData
	SetMessage(v string) *ListAsrVocabResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAsrVocabResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAsrVocabResponseBody
	GetSuccess() *bool
}

type ListAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAsrVocabResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 66E1ACB8-17B2-4BE8-8581-954A8EE1324B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAsrVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAsrVocabResponseBody) GetData() *ListAsrVocabResponseBodyData {
	return s.Data
}

func (s *ListAsrVocabResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAsrVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsrVocabResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAsrVocabResponseBody) SetCode(v string) *ListAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetData(v *ListAsrVocabResponseBodyData) *ListAsrVocabResponseBody {
	s.Data = v
	return s
}

func (s *ListAsrVocabResponseBody) SetMessage(v string) *ListAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetRequestId(v string) *ListAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsrVocabResponseBody) SetSuccess(v bool) *ListAsrVocabResponseBody {
	s.Success = &v
	return s
}

func (s *ListAsrVocabResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAsrVocabResponseBodyData struct {
	AsrVocab []*ListAsrVocabResponseBodyDataAsrVocab `json:"AsrVocab,omitempty" xml:"AsrVocab,omitempty" type:"Repeated"`
}

func (s ListAsrVocabResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAsrVocabResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBodyData) GetAsrVocab() []*ListAsrVocabResponseBodyDataAsrVocab {
	return s.AsrVocab
}

func (s *ListAsrVocabResponseBodyData) SetAsrVocab(v []*ListAsrVocabResponseBodyDataAsrVocab) *ListAsrVocabResponseBodyData {
	s.AsrVocab = v
	return s
}

func (s *ListAsrVocabResponseBodyData) Validate() error {
	if s.AsrVocab != nil {
		for _, item := range s.AsrVocab {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAsrVocabResponseBodyDataAsrVocab struct {
	AsrVersion *int32 `json:"AsrVersion,omitempty" xml:"AsrVersion,omitempty"`
	// example:
	//
	// 2019-04-15T14:57Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 18
	Id                   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ModelCustomizationId *string `json:"ModelCustomizationId,omitempty" xml:"ModelCustomizationId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2019-04-15T14:57Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// a01daaaxxxxxxxxx
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s ListAsrVocabResponseBodyDataAsrVocab) String() string {
	return dara.Prettify(s)
}

func (s ListAsrVocabResponseBodyDataAsrVocab) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetAsrVersion() *int32 {
	return s.AsrVersion
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetId() *string {
	return s.Id
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetModelCustomizationId() *string {
	return s.ModelCustomizationId
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetName() *string {
	return s.Name
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetAsrVersion(v int32) *ListAsrVocabResponseBodyDataAsrVocab {
	s.AsrVersion = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetCreateTime(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.CreateTime = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.Id = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetModelCustomizationId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.ModelCustomizationId = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetName(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.Name = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetUpdateTime(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.UpdateTime = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) SetVocabularyId(v string) *ListAsrVocabResponseBodyDataAsrVocab {
	s.VocabularyId = &v
	return s
}

func (s *ListAsrVocabResponseBodyDataAsrVocab) Validate() error {
	return dara.Validate(s)
}
