// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceEnginesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVoiceEnginesResponseBody
	GetCode() *string
	SetData(v *ListVoiceEnginesResponseBodyData) *ListVoiceEnginesResponseBody
	GetData() *ListVoiceEnginesResponseBodyData
	SetHttpStatusCode(v int32) *ListVoiceEnginesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVoiceEnginesResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListVoiceEnginesResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListVoiceEnginesResponseBody
	GetRequestId() *string
}

type ListVoiceEnginesResponseBody struct {
	// The internal error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paginated results of voice engines.
	Data *ListVoiceEnginesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of dynamic error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVoiceEnginesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceEnginesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVoiceEnginesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVoiceEnginesResponseBody) GetData() *ListVoiceEnginesResponseBodyData {
	return s.Data
}

func (s *ListVoiceEnginesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVoiceEnginesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVoiceEnginesResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListVoiceEnginesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVoiceEnginesResponseBody) SetCode(v string) *ListVoiceEnginesResponseBody {
	s.Code = &v
	return s
}

func (s *ListVoiceEnginesResponseBody) SetData(v *ListVoiceEnginesResponseBodyData) *ListVoiceEnginesResponseBody {
	s.Data = v
	return s
}

func (s *ListVoiceEnginesResponseBody) SetHttpStatusCode(v int32) *ListVoiceEnginesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVoiceEnginesResponseBody) SetMessage(v string) *ListVoiceEnginesResponseBody {
	s.Message = &v
	return s
}

func (s *ListVoiceEnginesResponseBody) SetParams(v []*string) *ListVoiceEnginesResponseBody {
	s.Params = v
	return s
}

func (s *ListVoiceEnginesResponseBody) SetRequestId(v string) *ListVoiceEnginesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVoiceEnginesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVoiceEnginesResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of voice engine objects.
	VoiceEngines []*ListVoiceEnginesResponseBodyDataVoiceEngines `json:"VoiceEngines,omitempty" xml:"VoiceEngines,omitempty" type:"Repeated"`
}

func (s ListVoiceEnginesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceEnginesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVoiceEnginesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVoiceEnginesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVoiceEnginesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVoiceEnginesResponseBodyData) GetVoiceEngines() []*ListVoiceEnginesResponseBodyDataVoiceEngines {
	return s.VoiceEngines
}

func (s *ListVoiceEnginesResponseBodyData) SetPageNumber(v int32) *ListVoiceEnginesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListVoiceEnginesResponseBodyData) SetPageSize(v int32) *ListVoiceEnginesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListVoiceEnginesResponseBodyData) SetTotalCount(v int32) *ListVoiceEnginesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListVoiceEnginesResponseBodyData) SetVoiceEngines(v []*ListVoiceEnginesResponseBodyDataVoiceEngines) *ListVoiceEnginesResponseBodyData {
	s.VoiceEngines = v
	return s
}

func (s *ListVoiceEnginesResponseBodyData) Validate() error {
	if s.VoiceEngines != nil {
		for _, item := range s.VoiceEngines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVoiceEnginesResponseBodyDataVoiceEngines struct {
	// The engine configuration.
	//
	// example:
	//
	// [{\\"displayName\\":\\"AppKey\\",\\"maxLength\\":64,\\"name\\":\\"AppKey\\",\\"order\\":1,\\"required\\":true},{\\"displayName\\":\\"AccessKey\\",\\"maxLength\\":64,\\"name\\":\\"AccessKey\\",\\"order\\":2,\\"required\\":true}]
	ConfigSchema *string `json:"ConfigSchema,omitempty" xml:"ConfigSchema,omitempty"`
	// The voice engine.
	//
	// example:
	//
	// VOLC
	NlsEngine *string `json:"NlsEngine,omitempty" xml:"NlsEngine,omitempty"`
	// The display name.
	//
	// example:
	//
	// Doubao
	NlsEngineName *string `json:"NlsEngineName,omitempty" xml:"NlsEngineName,omitempty"`
}

func (s ListVoiceEnginesResponseBodyDataVoiceEngines) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceEnginesResponseBodyDataVoiceEngines) GoString() string {
	return s.String()
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) GetConfigSchema() *string {
	return s.ConfigSchema
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) GetNlsEngine() *string {
	return s.NlsEngine
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) GetNlsEngineName() *string {
	return s.NlsEngineName
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) SetConfigSchema(v string) *ListVoiceEnginesResponseBodyDataVoiceEngines {
	s.ConfigSchema = &v
	return s
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) SetNlsEngine(v string) *ListVoiceEnginesResponseBodyDataVoiceEngines {
	s.NlsEngine = &v
	return s
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) SetNlsEngineName(v string) *ListVoiceEnginesResponseBodyDataVoiceEngines {
	s.NlsEngineName = &v
	return s
}

func (s *ListVoiceEnginesResponseBodyDataVoiceEngines) Validate() error {
	return dara.Validate(s)
}
