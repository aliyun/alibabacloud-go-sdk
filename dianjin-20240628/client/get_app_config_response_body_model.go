// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetAppConfigResponseBody
	GetCost() *int64
	SetData(v *GetAppConfigResponseBodyData) *GetAppConfigResponseBody
	GetData() *GetAppConfigResponseBodyData
	SetDataType(v string) *GetAppConfigResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetAppConfigResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetAppConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppConfigResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetAppConfigResponseBody
	GetTime() *string
}

type GetAppConfigResponseBody struct {
	// example:
	//
	// null
	Cost *int64                        `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetAppConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetAppConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetAppConfigResponseBody) GetData() *GetAppConfigResponseBodyData {
	return s.Data
}

func (s *GetAppConfigResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetAppConfigResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetAppConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppConfigResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetAppConfigResponseBody) SetCost(v int64) *GetAppConfigResponseBody {
	s.Cost = &v
	return s
}

func (s *GetAppConfigResponseBody) SetData(v *GetAppConfigResponseBodyData) *GetAppConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAppConfigResponseBody) SetDataType(v string) *GetAppConfigResponseBody {
	s.DataType = &v
	return s
}

func (s *GetAppConfigResponseBody) SetErrCode(v string) *GetAppConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAppConfigResponseBody) SetMessage(v string) *GetAppConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppConfigResponseBody) SetRequestId(v string) *GetAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppConfigResponseBody) SetSuccess(v bool) *GetAppConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppConfigResponseBody) SetTime(v string) *GetAppConfigResponseBody {
	s.Time = &v
	return s
}

func (s *GetAppConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppConfigResponseBodyData struct {
	EmbeddingTypeList         []map[string]*string `json:"embeddingTypeList,omitempty" xml:"embeddingTypeList,omitempty" type:"Repeated"`
	FrontendConfig            map[string]*bool     `json:"frontendConfig,omitempty" xml:"frontendConfig,omitempty"`
	LibraryDocumentStatusList []map[string]*string `json:"libraryDocumentStatusList,omitempty" xml:"libraryDocumentStatusList,omitempty" type:"Repeated"`
	LlmHelperTypeList         []map[string]*string `json:"llmHelperTypeList,omitempty" xml:"llmHelperTypeList,omitempty" type:"Repeated"`
	TextIndexCategoryList     []*string            `json:"textIndexCategoryList,omitempty" xml:"textIndexCategoryList,omitempty" type:"Repeated"`
	VectorIndexCategoryList   []*string            `json:"vectorIndexCategoryList,omitempty" xml:"vectorIndexCategoryList,omitempty" type:"Repeated"`
}

func (s GetAppConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBodyData) GetEmbeddingTypeList() []map[string]*string {
	return s.EmbeddingTypeList
}

func (s *GetAppConfigResponseBodyData) GetFrontendConfig() map[string]*bool {
	return s.FrontendConfig
}

func (s *GetAppConfigResponseBodyData) GetLibraryDocumentStatusList() []map[string]*string {
	return s.LibraryDocumentStatusList
}

func (s *GetAppConfigResponseBodyData) GetLlmHelperTypeList() []map[string]*string {
	return s.LlmHelperTypeList
}

func (s *GetAppConfigResponseBodyData) GetTextIndexCategoryList() []*string {
	return s.TextIndexCategoryList
}

func (s *GetAppConfigResponseBodyData) GetVectorIndexCategoryList() []*string {
	return s.VectorIndexCategoryList
}

func (s *GetAppConfigResponseBodyData) SetEmbeddingTypeList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.EmbeddingTypeList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetFrontendConfig(v map[string]*bool) *GetAppConfigResponseBodyData {
	s.FrontendConfig = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetLibraryDocumentStatusList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.LibraryDocumentStatusList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetLlmHelperTypeList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.LlmHelperTypeList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetTextIndexCategoryList(v []*string) *GetAppConfigResponseBodyData {
	s.TextIndexCategoryList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetVectorIndexCategoryList(v []*string) *GetAppConfigResponseBodyData {
	s.VectorIndexCategoryList = v
	return s
}

func (s *GetAppConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
