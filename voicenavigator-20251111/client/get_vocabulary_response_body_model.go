// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabularyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVocabularyResponseBody
	GetCode() *string
	SetData(v *GetVocabularyResponseBodyData) *GetVocabularyResponseBody
	GetData() *GetVocabularyResponseBodyData
	SetHttpStatusCode(v int32) *GetVocabularyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVocabularyResponseBody
	GetMessage() *string
	SetParams(v []*string) *GetVocabularyResponseBody
	GetParams() []*string
	SetRequestId(v string) *GetVocabularyResponseBody
	GetRequestId() *string
}

type GetVocabularyResponseBody struct {
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetVocabularyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 97E7ED13-6884-52D7-B97E-C780D7870680
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVocabularyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVocabularyResponseBody) GoString() string {
	return s.String()
}

func (s *GetVocabularyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVocabularyResponseBody) GetData() *GetVocabularyResponseBodyData {
	return s.Data
}

func (s *GetVocabularyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVocabularyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVocabularyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GetVocabularyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVocabularyResponseBody) SetCode(v string) *GetVocabularyResponseBody {
	s.Code = &v
	return s
}

func (s *GetVocabularyResponseBody) SetData(v *GetVocabularyResponseBodyData) *GetVocabularyResponseBody {
	s.Data = v
	return s
}

func (s *GetVocabularyResponseBody) SetHttpStatusCode(v int32) *GetVocabularyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVocabularyResponseBody) SetMessage(v string) *GetVocabularyResponseBody {
	s.Message = &v
	return s
}

func (s *GetVocabularyResponseBody) SetParams(v []*string) *GetVocabularyResponseBody {
	s.Params = v
	return s
}

func (s *GetVocabularyResponseBody) SetRequestId(v string) *GetVocabularyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVocabularyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVocabularyResponseBodyData struct {
	// example:
	//
	// 1773453676000
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
	// 1773453676000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	// example:
	//
	// 50
	WordCount *string `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	Words     *string `json:"Words,omitempty" xml:"Words,omitempty"`
}

func (s GetVocabularyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVocabularyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVocabularyResponseBodyData) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetVocabularyResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetVocabularyResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVocabularyResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetVocabularyResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetVocabularyResponseBodyData) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetVocabularyResponseBodyData) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *GetVocabularyResponseBodyData) GetWordCount() *string {
	return s.WordCount
}

func (s *GetVocabularyResponseBodyData) GetWords() *string {
	return s.Words
}

func (s *GetVocabularyResponseBodyData) SetCreatedTime(v int64) *GetVocabularyResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetDescription(v string) *GetVocabularyResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetInstanceId(v string) *GetVocabularyResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetName(v string) *GetVocabularyResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetTenantId(v string) *GetVocabularyResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetUpdatedTime(v int64) *GetVocabularyResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetVocabularyId(v string) *GetVocabularyResponseBodyData {
	s.VocabularyId = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetWordCount(v string) *GetVocabularyResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *GetVocabularyResponseBodyData) SetWords(v string) *GetVocabularyResponseBodyData {
	s.Words = &v
	return s
}

func (s *GetVocabularyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
