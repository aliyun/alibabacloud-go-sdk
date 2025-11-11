// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDatasetsResponseBody
	GetCode() *string
	SetCustomSemanticSearchConfig(v *ListDatasetsResponseBodyCustomSemanticSearchConfig) *ListDatasetsResponseBody
	GetCustomSemanticSearchConfig() *ListDatasetsResponseBodyCustomSemanticSearchConfig
	SetData(v []*ListDatasetsResponseBodyData) *ListDatasetsResponseBody
	GetData() []*ListDatasetsResponseBodyData
	SetHttpStatusCode(v int32) *ListDatasetsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDatasetsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListDatasetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDatasetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDatasetsResponseBody
	GetSuccess() *bool
	SetThirdSearchConfig(v *ListDatasetsResponseBodyThirdSearchConfig) *ListDatasetsResponseBody
	GetThirdSearchConfig() *ListDatasetsResponseBodyThirdSearchConfig
	SetTotalCount(v int32) *ListDatasetsResponseBody
	GetTotalCount() *int32
}

type ListDatasetsResponseBody struct {
	// example:
	//
	// NoData
	Code                       *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	CustomSemanticSearchConfig *ListDatasetsResponseBodyCustomSemanticSearchConfig `json:"CustomSemanticSearchConfig,omitempty" xml:"CustomSemanticSearchConfig,omitempty" type:"Struct"`
	Data                       []*ListDatasetsResponseBodyData                     `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success           *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	ThirdSearchConfig *ListDatasetsResponseBodyThirdSearchConfig `json:"ThirdSearchConfig,omitempty" xml:"ThirdSearchConfig,omitempty" type:"Struct"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDatasetsResponseBody) GetCustomSemanticSearchConfig() *ListDatasetsResponseBodyCustomSemanticSearchConfig {
	return s.CustomSemanticSearchConfig
}

func (s *ListDatasetsResponseBody) GetData() []*ListDatasetsResponseBodyData {
	return s.Data
}

func (s *ListDatasetsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDatasetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDatasetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDatasetsResponseBody) GetThirdSearchConfig() *ListDatasetsResponseBodyThirdSearchConfig {
	return s.ThirdSearchConfig
}

func (s *ListDatasetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDatasetsResponseBody) SetCode(v string) *ListDatasetsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatasetsResponseBody) SetCustomSemanticSearchConfig(v *ListDatasetsResponseBodyCustomSemanticSearchConfig) *ListDatasetsResponseBody {
	s.CustomSemanticSearchConfig = v
	return s
}

func (s *ListDatasetsResponseBody) SetData(v []*ListDatasetsResponseBodyData) *ListDatasetsResponseBody {
	s.Data = v
	return s
}

func (s *ListDatasetsResponseBody) SetHttpStatusCode(v int32) *ListDatasetsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDatasetsResponseBody) SetMessage(v string) *ListDatasetsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatasetsResponseBody) SetPageNumber(v int32) *ListDatasetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsResponseBody) SetPageSize(v int32) *ListDatasetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) SetSuccess(v bool) *ListDatasetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDatasetsResponseBody) SetThirdSearchConfig(v *ListDatasetsResponseBodyThirdSearchConfig) *ListDatasetsResponseBody {
	s.ThirdSearchConfig = v
	return s
}

func (s *ListDatasetsResponseBody) SetTotalCount(v int32) *ListDatasetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetsResponseBody) Validate() error {
	if s.CustomSemanticSearchConfig != nil {
		if err := s.CustomSemanticSearchConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ThirdSearchConfig != nil {
		if err := s.ThirdSearchConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDatasetsResponseBodyCustomSemanticSearchConfig struct {
	// example:
	//
	// 3
	DatasetQuota *int32 `json:"DatasetQuota,omitempty" xml:"DatasetQuota,omitempty"`
	// example:
	//
	// 1
	DatasetUsedQuota *int32 `json:"DatasetUsedQuota,omitempty" xml:"DatasetUsedQuota,omitempty"`
	// example:
	//
	// 1000
	DocQuota *int64 `json:"DocQuota,omitempty" xml:"DocQuota,omitempty"`
	// example:
	//
	// 1
	DocUsedQuota *int64 `json:"DocUsedQuota,omitempty" xml:"DocUsedQuota,omitempty"`
}

func (s ListDatasetsResponseBodyCustomSemanticSearchConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyCustomSemanticSearchConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) GetDatasetQuota() *int32 {
	return s.DatasetQuota
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) GetDatasetUsedQuota() *int32 {
	return s.DatasetUsedQuota
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) GetDocQuota() *int64 {
	return s.DocQuota
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) GetDocUsedQuota() *int64 {
	return s.DocUsedQuota
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) SetDatasetQuota(v int32) *ListDatasetsResponseBodyCustomSemanticSearchConfig {
	s.DatasetQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) SetDatasetUsedQuota(v int32) *ListDatasetsResponseBodyCustomSemanticSearchConfig {
	s.DatasetUsedQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) SetDocQuota(v int64) *ListDatasetsResponseBodyCustomSemanticSearchConfig {
	s.DocQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) SetDocUsedQuota(v int64) *ListDatasetsResponseBodyCustomSemanticSearchConfig {
	s.DocUsedQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyCustomSemanticSearchConfig) Validate() error {
	return dara.Validate(s)
}

type ListDatasetsResponseBodyData struct {
	// example:
	//
	// 2024-11-12 21:46:24
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// xxx
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// 1
	DatasetId *int64 `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// xxx
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// example:
	//
	// CustomSemanticSearch
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// example:
	//
	// 1
	DocUsedQuota *int64 `json:"DocUsedQuota,omitempty" xml:"DocUsedQuota,omitempty"`
	// example:
	//
	// 1
	SearchDatasetEnable *int32 `json:"SearchDatasetEnable,omitempty" xml:"SearchDatasetEnable,omitempty"`
}

func (s ListDatasetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDatasetsResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDatasetsResponseBodyData) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *ListDatasetsResponseBodyData) GetDatasetId() *int64 {
	return s.DatasetId
}

func (s *ListDatasetsResponseBodyData) GetDatasetName() *string {
	return s.DatasetName
}

func (s *ListDatasetsResponseBodyData) GetDatasetType() *string {
	return s.DatasetType
}

func (s *ListDatasetsResponseBodyData) GetDocUsedQuota() *int64 {
	return s.DocUsedQuota
}

func (s *ListDatasetsResponseBodyData) GetSearchDatasetEnable() *int32 {
	return s.SearchDatasetEnable
}

func (s *ListDatasetsResponseBodyData) SetCreateTime(v string) *ListDatasetsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetCreateUser(v string) *ListDatasetsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetDatasetDescription(v string) *ListDatasetsResponseBodyData {
	s.DatasetDescription = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetDatasetId(v int64) *ListDatasetsResponseBodyData {
	s.DatasetId = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetDatasetName(v string) *ListDatasetsResponseBodyData {
	s.DatasetName = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetDatasetType(v string) *ListDatasetsResponseBodyData {
	s.DatasetType = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetDocUsedQuota(v int64) *ListDatasetsResponseBodyData {
	s.DocUsedQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyData) SetSearchDatasetEnable(v int32) *ListDatasetsResponseBodyData {
	s.SearchDatasetEnable = &v
	return s
}

func (s *ListDatasetsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDatasetsResponseBodyThirdSearchConfig struct {
	// example:
	//
	// 2
	DatasetQuota *int32 `json:"DatasetQuota,omitempty" xml:"DatasetQuota,omitempty"`
	// example:
	//
	// 1
	DatasetUsedQuota *int32 `json:"DatasetUsedQuota,omitempty" xml:"DatasetUsedQuota,omitempty"`
}

func (s ListDatasetsResponseBodyThirdSearchConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBodyThirdSearchConfig) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBodyThirdSearchConfig) GetDatasetQuota() *int32 {
	return s.DatasetQuota
}

func (s *ListDatasetsResponseBodyThirdSearchConfig) GetDatasetUsedQuota() *int32 {
	return s.DatasetUsedQuota
}

func (s *ListDatasetsResponseBodyThirdSearchConfig) SetDatasetQuota(v int32) *ListDatasetsResponseBodyThirdSearchConfig {
	s.DatasetQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyThirdSearchConfig) SetDatasetUsedQuota(v int32) *ListDatasetsResponseBodyThirdSearchConfig {
	s.DatasetUsedQuota = &v
	return s
}

func (s *ListDatasetsResponseBodyThirdSearchConfig) Validate() error {
	return dara.Validate(s)
}
