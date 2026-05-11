// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAsrConfigResponseBody
	GetCode() *string
	SetData(v *GetAsrConfigResponseBodyData) *GetAsrConfigResponseBody
	GetData() *GetAsrConfigResponseBodyData
	SetErrorMsg(v string) *GetAsrConfigResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GetAsrConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetAsrConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAsrConfigResponseBody
	GetSuccess() *bool
}

type GetAsrConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAsrConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Not Found
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAsrConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsrConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsrConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAsrConfigResponseBody) GetData() *GetAsrConfigResponseBodyData {
	return s.Data
}

func (s *GetAsrConfigResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetAsrConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAsrConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsrConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAsrConfigResponseBody) SetCode(v string) *GetAsrConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetData(v *GetAsrConfigResponseBodyData) *GetAsrConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAsrConfigResponseBody) SetErrorMsg(v string) *GetAsrConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetHttpStatusCode(v int32) *GetAsrConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetRequestId(v string) *GetAsrConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsrConfigResponseBody) SetSuccess(v bool) *GetAsrConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsrConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAsrConfigResponseBodyData struct {
	// example:
	//
	// your-app-key
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrAcousticModelId *string `json:"AsrAcousticModelId,omitempty" xml:"AsrAcousticModelId,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrClassVocabularyId *string `json:"AsrClassVocabularyId,omitempty" xml:"AsrClassVocabularyId,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrCustomizationId *string `json:"AsrCustomizationId,omitempty" xml:"AsrCustomizationId,omitempty"`
	AsrOverrides       *string `json:"AsrOverrides,omitempty" xml:"AsrOverrides,omitempty"`
	// example:
	//
	// 3b1d3031-8b6e-460a-8640-d330f2ca50b8
	AsrVocabularyId *string `json:"AsrVocabularyId,omitempty" xml:"AsrVocabularyId,omitempty"`
	// example:
	//
	// ali
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// EngineXufei
	EngineXufei *string `json:"EngineXufei,omitempty" xml:"EngineXufei,omitempty"`
	// example:
	//
	// Authorized
	NlsServiceType *string `json:"NlsServiceType,omitempty" xml:"NlsServiceType,omitempty"`
}

func (s GetAsrConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAsrConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsrConfigResponseBodyData) GetAppKey() *string {
	return s.AppKey
}

func (s *GetAsrConfigResponseBodyData) GetAsrAcousticModelId() *string {
	return s.AsrAcousticModelId
}

func (s *GetAsrConfigResponseBodyData) GetAsrClassVocabularyId() *string {
	return s.AsrClassVocabularyId
}

func (s *GetAsrConfigResponseBodyData) GetAsrCustomizationId() *string {
	return s.AsrCustomizationId
}

func (s *GetAsrConfigResponseBodyData) GetAsrOverrides() *string {
	return s.AsrOverrides
}

func (s *GetAsrConfigResponseBodyData) GetAsrVocabularyId() *string {
	return s.AsrVocabularyId
}

func (s *GetAsrConfigResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *GetAsrConfigResponseBodyData) GetEngineXufei() *string {
	return s.EngineXufei
}

func (s *GetAsrConfigResponseBodyData) GetNlsServiceType() *string {
	return s.NlsServiceType
}

func (s *GetAsrConfigResponseBodyData) SetAppKey(v string) *GetAsrConfigResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrAcousticModelId(v string) *GetAsrConfigResponseBodyData {
	s.AsrAcousticModelId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrClassVocabularyId(v string) *GetAsrConfigResponseBodyData {
	s.AsrClassVocabularyId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrCustomizationId(v string) *GetAsrConfigResponseBodyData {
	s.AsrCustomizationId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrOverrides(v string) *GetAsrConfigResponseBodyData {
	s.AsrOverrides = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetAsrVocabularyId(v string) *GetAsrConfigResponseBodyData {
	s.AsrVocabularyId = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetEngine(v string) *GetAsrConfigResponseBodyData {
	s.Engine = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetEngineXufei(v string) *GetAsrConfigResponseBodyData {
	s.EngineXufei = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) SetNlsServiceType(v string) *GetAsrConfigResponseBodyData {
	s.NlsServiceType = &v
	return s
}

func (s *GetAsrConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
