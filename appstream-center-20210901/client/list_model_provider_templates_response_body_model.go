// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListModelProviderTemplatesResponseBodyData) *ListModelProviderTemplatesResponseBody
	GetData() []*ListModelProviderTemplatesResponseBodyData
	SetPageNumber(v int32) *ListModelProviderTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProviderTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModelProviderTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModelProviderTemplatesResponseBody
	GetTotalCount() *int32
}

type ListModelProviderTemplatesResponseBody struct {
	Data []*ListModelProviderTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelProviderTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelProviderTemplatesResponseBody) GetData() []*ListModelProviderTemplatesResponseBodyData {
	return s.Data
}

func (s *ListModelProviderTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProviderTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProviderTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelProviderTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelProviderTemplatesResponseBody) SetData(v []*ListModelProviderTemplatesResponseBodyData) *ListModelProviderTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListModelProviderTemplatesResponseBody) SetPageNumber(v int32) *ListModelProviderTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBody) SetPageSize(v int32) *ListModelProviderTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBody) SetRequestId(v string) *ListModelProviderTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBody) SetTotalCount(v int32) *ListModelProviderTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelProviderTemplatesResponseBodyData struct {
	// example:
	//
	// {
	//
	// 	"api": "openai-completions",
	//
	// 	"apiKey": "sk-xxxx",
	//
	// 	"baseUrl": "https://dashscope.aliyuncs.com/compatible-mode/v1"
	//
	// }
	Config      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	EnableWuyingProxy *bool   `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
	// example:
	//
	// WuyingCredit
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
}

func (s ListModelProviderTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelProviderTemplatesResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *ListModelProviderTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListModelProviderTemplatesResponseBodyData) GetEnableWuyingProxy() *bool {
	return s.EnableWuyingProxy
}

func (s *ListModelProviderTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListModelProviderTemplatesResponseBodyData) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProviderTemplatesResponseBodyData) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ListModelProviderTemplatesResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListModelProviderTemplatesResponseBodyData) SetConfig(v string) *ListModelProviderTemplatesResponseBodyData {
	s.Config = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) SetDescription(v string) *ListModelProviderTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) SetEnableWuyingProxy(v bool) *ListModelProviderTemplatesResponseBodyData {
	s.EnableWuyingProxy = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) SetName(v string) *ListModelProviderTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) SetProviderName(v string) *ListModelProviderTemplatesResponseBodyData {
	s.ProviderName = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) SetProviderTemplateId(v string) *ListModelProviderTemplatesResponseBodyData {
	s.ProviderTemplateId = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) SetProviderType(v string) *ListModelProviderTemplatesResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *ListModelProviderTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
