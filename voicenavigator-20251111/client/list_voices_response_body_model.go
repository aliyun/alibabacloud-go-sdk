// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoicesResponseBody
	GetCode() *string
	SetData(v *ListVoicesResponseBodyData) *ListVoicesResponseBody
	GetData() *ListVoicesResponseBodyData
	SetHttpStatusCode(v int32) *ListVoicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVoicesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListVoicesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListVoicesResponseBody
	GetRequestId() *string
}

type ListVoicesResponseBody struct {
	// example:
	//
	// OK
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListVoicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListVoicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoicesResponseBody) GetData() *ListVoicesResponseBodyData {
	return s.Data
}

func (s *ListVoicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVoicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVoicesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListVoicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoicesResponseBody) SetCode(v string) *ListVoicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoicesResponseBody) SetData(v *ListVoicesResponseBodyData) *ListVoicesResponseBody {
	s.Data = v
	return s
}

func (s *ListVoicesResponseBody) SetHttpStatusCode(v int32) *ListVoicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVoicesResponseBody) SetMessage(v string) *ListVoicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoicesResponseBody) SetParams(v []*string) *ListVoicesResponseBody {
	s.Params = v
	return s
}

func (s *ListVoicesResponseBody) SetRequestId(v string) *ListVoicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoicesResponseBodyData struct {
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
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Voices     []*ListVoicesResponseBodyDataVoices `json:"Voices,omitempty" xml:"Voices,omitempty" type:"Repeated"`
}

func (s ListVoicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVoicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVoicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVoicesResponseBodyData) GetVoices() []*ListVoicesResponseBodyDataVoices {
	return s.Voices
}

func (s *ListVoicesResponseBodyData) SetPageNumber(v int32) *ListVoicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVoicesResponseBodyData) SetPageSize(v int32) *ListVoicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVoicesResponseBodyData) SetTotalCount(v int32) *ListVoicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVoicesResponseBodyData) SetVoices(v []*ListVoicesResponseBodyDataVoices) *ListVoicesResponseBodyData {
	s.Voices = v
	return s
}

func (s *ListVoicesResponseBodyData) Validate() error {
	if s.Voices != nil {
		for _, item := range s.Voices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVoicesResponseBodyDataVoices struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// Qwen
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// BAILIAN
	NlsEngine       *string   `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	Style           *string   `json:"Style,omitempty" xml:"Style,omitempty"`
	SupportedParams []*string `json:"SupportedParams,omitempty" xml:"SupportedParams,omitempty" type:"Repeated"`
	// example:
	//
	// Cherry
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
}

func (s ListVoicesResponseBodyDataVoices) String() string {
	return dara.Prettify(s)
}

func (s ListVoicesResponseBodyDataVoices) GoString() string {
	return s.String()
}

func (s *ListVoicesResponseBodyDataVoices) GetCategory() *string {
	return s.Category
}

func (s *ListVoicesResponseBodyDataVoices) GetLanguage() *string {
	return s.Language
}

func (s *ListVoicesResponseBodyDataVoices) GetModel() *string {
	return s.Model
}

func (s *ListVoicesResponseBodyDataVoices) GetName() *string {
	return s.Name
}

func (s *ListVoicesResponseBodyDataVoices) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListVoicesResponseBodyDataVoices) GetStyle() *string {
	return s.Style
}

func (s *ListVoicesResponseBodyDataVoices) GetSupportedParams() []*string {
	return s.SupportedParams
}

func (s *ListVoicesResponseBodyDataVoices) GetVoice() *string {
	return s.Voice
}

func (s *ListVoicesResponseBodyDataVoices) SetCategory(v string) *ListVoicesResponseBodyDataVoices {
	s.Category = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetLanguage(v string) *ListVoicesResponseBodyDataVoices {
	s.Language = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetModel(v string) *ListVoicesResponseBodyDataVoices {
	s.Model = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetName(v string) *ListVoicesResponseBodyDataVoices {
	s.Name = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetNlsEngine(v string) *ListVoicesResponseBodyDataVoices {
	s.NlsEngine = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetStyle(v string) *ListVoicesResponseBodyDataVoices {
	s.Style = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetSupportedParams(v []*string) *ListVoicesResponseBodyDataVoices {
	s.SupportedParams = v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) SetVoice(v string) *ListVoicesResponseBodyDataVoices {
	s.Voice = &v
	return s
}

func (s *ListVoicesResponseBodyDataVoices) Validate() error {
	return dara.Validate(s)
}
