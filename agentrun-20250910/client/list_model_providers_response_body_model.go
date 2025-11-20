// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelProvidersResponseBody
	GetCode() *string
	SetData(v *ListModelProvidersResponseBodyData) *ListModelProvidersResponseBody
	GetData() *ListModelProvidersResponseBodyData
	SetRequestId(v string) *ListModelProvidersResponseBody
	GetRequestId() *string
}

type ListModelProvidersResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListModelProvidersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelProvidersResponseBody) GetData() *ListModelProvidersResponseBodyData {
	return s.Data
}

func (s *ListModelProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelProvidersResponseBody) SetCode(v string) *ListModelProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelProvidersResponseBody) SetData(v *ListModelProvidersResponseBodyData) *ListModelProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ListModelProvidersResponseBody) SetRequestId(v string) *ListModelProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelProvidersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelProvidersResponseBodyData struct {
	Items []*ListModelProvidersResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 211
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListModelProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelProvidersResponseBodyData) GetItems() []*ListModelProvidersResponseBodyDataItems {
	return s.Items
}

func (s *ListModelProvidersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProvidersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProvidersResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListModelProvidersResponseBodyData) SetItems(v []*ListModelProvidersResponseBodyDataItems) *ListModelProvidersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListModelProvidersResponseBodyData) SetPageNumber(v int32) *ListModelProvidersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListModelProvidersResponseBodyData) SetPageSize(v int32) *ListModelProvidersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListModelProvidersResponseBodyData) SetTotal(v int64) *ListModelProvidersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListModelProvidersResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelProvidersResponseBodyDataItems struct {
	// baseUrl
	//
	// example:
	//
	// baseUrl
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// modelInfoConfig
	ModelInfoConfig *ModelInfoConfig `json:"modelInfoConfig,omitempty" xml:"modelInfoConfig,omitempty"`
	// example:
	//
	// native
	ModelType *string   `json:"modelType,omitempty" xml:"modelType,omitempty"`
	Models    []*string `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// example:
	//
	// providerName
	ProviderName *string `json:"providerName,omitempty" xml:"providerName,omitempty"`
}

func (s ListModelProvidersResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListModelProvidersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListModelProvidersResponseBodyDataItems) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ListModelProvidersResponseBodyDataItems) GetModelInfoConfig() *ModelInfoConfig {
	return s.ModelInfoConfig
}

func (s *ListModelProvidersResponseBodyDataItems) GetModelType() *string {
	return s.ModelType
}

func (s *ListModelProvidersResponseBodyDataItems) GetModels() []*string {
	return s.Models
}

func (s *ListModelProvidersResponseBodyDataItems) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProvidersResponseBodyDataItems) SetBaseUrl(v string) *ListModelProvidersResponseBodyDataItems {
	s.BaseUrl = &v
	return s
}

func (s *ListModelProvidersResponseBodyDataItems) SetModelInfoConfig(v *ModelInfoConfig) *ListModelProvidersResponseBodyDataItems {
	s.ModelInfoConfig = v
	return s
}

func (s *ListModelProvidersResponseBodyDataItems) SetModelType(v string) *ListModelProvidersResponseBodyDataItems {
	s.ModelType = &v
	return s
}

func (s *ListModelProvidersResponseBodyDataItems) SetModels(v []*string) *ListModelProvidersResponseBodyDataItems {
	s.Models = v
	return s
}

func (s *ListModelProvidersResponseBodyDataItems) SetProviderName(v string) *ListModelProvidersResponseBodyDataItems {
	s.ProviderName = &v
	return s
}

func (s *ListModelProvidersResponseBodyDataItems) Validate() error {
	if s.ModelInfoConfig != nil {
		if err := s.ModelInfoConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
