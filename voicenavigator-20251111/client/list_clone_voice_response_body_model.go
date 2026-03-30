// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCloneVoiceResponseBody
	GetCode() *string
	SetData(v *ListCloneVoiceResponseBodyData) *ListCloneVoiceResponseBody
	GetData() *ListCloneVoiceResponseBodyData
	SetHttpStatusCode(v int32) *ListCloneVoiceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCloneVoiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCloneVoiceResponseBody
	GetRequestId() *string
}

type ListCloneVoiceResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCloneVoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCloneVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCloneVoiceResponseBody) GetData() *ListCloneVoiceResponseBodyData {
	return s.Data
}

func (s *ListCloneVoiceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCloneVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCloneVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloneVoiceResponseBody) SetCode(v string) *ListCloneVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloneVoiceResponseBody) SetData(v *ListCloneVoiceResponseBodyData) *ListCloneVoiceResponseBody {
	s.Data = v
	return s
}

func (s *ListCloneVoiceResponseBody) SetHttpStatusCode(v int32) *ListCloneVoiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCloneVoiceResponseBody) SetMessage(v string) *ListCloneVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloneVoiceResponseBody) SetRequestId(v string) *ListCloneVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloneVoiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloneVoiceResponseBodyData struct {
	CloneVoices []*ListCloneVoiceResponseBodyDataCloneVoices `json:"CloneVoices,omitempty" xml:"CloneVoices,omitempty" type:"Repeated"`
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
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloneVoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceResponseBodyData) GetCloneVoices() []*ListCloneVoiceResponseBodyDataCloneVoices {
	return s.CloneVoices
}

func (s *ListCloneVoiceResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloneVoiceResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloneVoiceResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloneVoiceResponseBodyData) SetCloneVoices(v []*ListCloneVoiceResponseBodyDataCloneVoices) *ListCloneVoiceResponseBodyData {
	s.CloneVoices = v
	return s
}

func (s *ListCloneVoiceResponseBodyData) SetPageNumber(v int32) *ListCloneVoiceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCloneVoiceResponseBodyData) SetPageSize(v int32) *ListCloneVoiceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCloneVoiceResponseBodyData) SetTotalCount(v int32) *ListCloneVoiceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCloneVoiceResponseBodyData) Validate() error {
	if s.CloneVoices != nil {
		for _, item := range s.CloneVoices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloneVoiceResponseBodyDataCloneVoices struct {
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66b
	CloneVoiceId *string `json:"CloneVoiceId,omitempty" xml:"CloneVoiceId,omitempty"`
	// example:
	//
	// 1764900994000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// CosyVoice
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1308144684576655
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1764900994000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// cosyvoice-v3-plus-voicebot2-3666e4bbb2b94832ac4f4107b5804c34
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s ListCloneVoiceResponseBodyDataCloneVoices) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceResponseBodyDataCloneVoices) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetCloneVoiceId() *string {
	return s.CloneVoiceId
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetModel() *string {
	return s.Model
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetName() *string {
	return s.Name
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetStatus() *string {
	return s.Status
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetTenantId() *string {
	return s.TenantId
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) GetVoice() *string {
	return s.Voice
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetCloneVoiceId(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.CloneVoiceId = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetCreatedTime(v int64) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.CreatedTime = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetInstanceId(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.InstanceId = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetModel(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.Model = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetName(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.Name = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetNlsEngine(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.NlsEngine = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetStatus(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.Status = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetTenantId(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.TenantId = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetUpdatedTime(v int64) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.UpdatedTime = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) SetVoice(v string) *ListCloneVoiceResponseBodyDataCloneVoices {
	s.Voice = &v
	return s
}

func (s *ListCloneVoiceResponseBodyDataCloneVoices) Validate() error {
	return dara.Validate(s)
}
