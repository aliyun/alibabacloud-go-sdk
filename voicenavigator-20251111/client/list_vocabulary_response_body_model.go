// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabularyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVocabularyResponseBody
	GetCode() *string
	SetData(v *ListVocabularyResponseBodyData) *ListVocabularyResponseBody
	GetData() *ListVocabularyResponseBodyData
	SetHttpStatusCode(v int32) *ListVocabularyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVocabularyResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListVocabularyResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListVocabularyResponseBody
	GetRequestId() *string
}

type ListVocabularyResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListVocabularyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9DB8BA95-4513-54B9-9C67-A28909CFB4AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVocabularyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVocabularyResponseBody) GoString() string {
	return s.String()
}

func (s *ListVocabularyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVocabularyResponseBody) GetData() *ListVocabularyResponseBodyData {
	return s.Data
}

func (s *ListVocabularyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVocabularyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVocabularyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListVocabularyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVocabularyResponseBody) SetCode(v string) *ListVocabularyResponseBody {
	s.Code = &v
	return s
}

func (s *ListVocabularyResponseBody) SetData(v *ListVocabularyResponseBodyData) *ListVocabularyResponseBody {
	s.Data = v
	return s
}

func (s *ListVocabularyResponseBody) SetHttpStatusCode(v int32) *ListVocabularyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVocabularyResponseBody) SetMessage(v string) *ListVocabularyResponseBody {
	s.Message = &v
	return s
}

func (s *ListVocabularyResponseBody) SetParams(v []*string) *ListVocabularyResponseBody {
	s.Params = v
	return s
}

func (s *ListVocabularyResponseBody) SetRequestId(v string) *ListVocabularyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVocabularyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVocabularyResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vocabularies []*ListVocabularyResponseBodyDataVocabularies `json:"Vocabularies,omitempty" xml:"Vocabularies,omitempty" type:"Repeated"`
}

func (s ListVocabularyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVocabularyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVocabularyResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVocabularyResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVocabularyResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVocabularyResponseBodyData) GetVocabularies() []*ListVocabularyResponseBodyDataVocabularies {
	return s.Vocabularies
}

func (s *ListVocabularyResponseBodyData) SetPageNumber(v int32) *ListVocabularyResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVocabularyResponseBodyData) SetPageSize(v int32) *ListVocabularyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVocabularyResponseBodyData) SetTotalCount(v int32) *ListVocabularyResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVocabularyResponseBodyData) SetVocabularies(v []*ListVocabularyResponseBodyDataVocabularies) *ListVocabularyResponseBodyData {
	s.Vocabularies = v
	return s
}

func (s *ListVocabularyResponseBodyData) Validate() error {
	if s.Vocabularies != nil {
		for _, item := range s.Vocabularies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVocabularyResponseBodyDataVocabularies struct {
	// example:
	//
	// 1754013825102
	CreatedTime *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1308144684576655
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1754013825102
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	// example:
	//
	// 50
	WordCount *int32 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s ListVocabularyResponseBodyDataVocabularies) String() string {
	return dara.Prettify(s)
}

func (s ListVocabularyResponseBodyDataVocabularies) GoString() string {
	return s.String()
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetDescription() *string {
	return s.Description
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetName() *string {
	return s.Name
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetTenantId() *string {
	return s.TenantId
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *ListVocabularyResponseBodyDataVocabularies) GetWordCount() *int32 {
	return s.WordCount
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetCreatedTime(v int64) *ListVocabularyResponseBodyDataVocabularies {
	s.CreatedTime = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetDescription(v string) *ListVocabularyResponseBodyDataVocabularies {
	s.Description = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetInstanceId(v string) *ListVocabularyResponseBodyDataVocabularies {
	s.InstanceId = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetName(v string) *ListVocabularyResponseBodyDataVocabularies {
	s.Name = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetTenantId(v string) *ListVocabularyResponseBodyDataVocabularies {
	s.TenantId = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetUpdatedTime(v int64) *ListVocabularyResponseBodyDataVocabularies {
	s.UpdatedTime = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetVocabularyId(v string) *ListVocabularyResponseBodyDataVocabularies {
	s.VocabularyId = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) SetWordCount(v int32) *ListVocabularyResponseBodyDataVocabularies {
	s.WordCount = &v
	return s
}

func (s *ListVocabularyResponseBodyDataVocabularies) Validate() error {
	return dara.Validate(s)
}
