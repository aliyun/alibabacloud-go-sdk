// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLlmTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListLlmTemplatesResponseBodyData) *ListLlmTemplatesResponseBody
	GetData() []*ListLlmTemplatesResponseBodyData
	SetPageNumber(v int32) *ListLlmTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLlmTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLlmTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLlmTemplatesResponseBody
	GetTotalCount() *int32
}

type ListLlmTemplatesResponseBody struct {
	// The list of returned data objects.
	Data []*ListLlmTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the current page.
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
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLlmTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBody) GetData() []*ListLlmTemplatesResponseBodyData {
	return s.Data
}

func (s *ListLlmTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLlmTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLlmTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLlmTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLlmTemplatesResponseBody) SetData(v []*ListLlmTemplatesResponseBodyData) *ListLlmTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetPageNumber(v int32) *ListLlmTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetPageSize(v int32) *ListLlmTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetRequestId(v string) *ListLlmTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) SetTotalCount(v int32) *ListLlmTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLlmTemplatesResponseBody) Validate() error {
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

type ListLlmTemplatesResponseBodyData struct {
	// The model configuration JSON object.
	//
	// example:
	//
	// {
	//
	// 	"id": "qwen3.6-plus",
	//
	// 	"cost": {
	//
	// 		"input": 0,
	//
	// 		"output": 0,
	//
	// 		"cacheRead": 0,
	//
	// 		"cacheWrite": 0
	//
	// 	},
	//
	// 	"name": "Qwen3.6-Plus",
	//
	// 	"input": ["image", "text"],
	//
	// 	"compat": {
	//
	// 		"supportsUsageInStreaming": true
	//
	// 	},
	//
	// 	"maxTokens": 65536,
	//
	// 	"reasoning": false,
	//
	// 	"contextWindow": 1000000
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The template description.
	//
	// example:
	//
	// 千问Plus系列模型
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether this is the default model in the associated model group.
	//
	// example:
	//
	// true
	IsDefaultModel *bool `json:"IsDefaultModel,omitempty" xml:"IsDefaultModel,omitempty"`
	// The model code.
	//
	// example:
	//
	// qwen3.6-plus
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// The model template ID.
	//
	// example:
	//
	// llmt-xxxx
	LlmTemplateId *string `json:"LlmTemplateId,omitempty" xml:"LlmTemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// Qwen3.6-Plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the model provider template.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s ListLlmTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLlmTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLlmTemplatesResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *ListLlmTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListLlmTemplatesResponseBodyData) GetIsDefaultModel() *bool {
	return s.IsDefaultModel
}

func (s *ListLlmTemplatesResponseBodyData) GetLlmCode() *string {
	return s.LlmCode
}

func (s *ListLlmTemplatesResponseBodyData) GetLlmTemplateId() *string {
	return s.LlmTemplateId
}

func (s *ListLlmTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListLlmTemplatesResponseBodyData) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *ListLlmTemplatesResponseBodyData) SetConfig(v string) *ListLlmTemplatesResponseBodyData {
	s.Config = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetDescription(v string) *ListLlmTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetIsDefaultModel(v bool) *ListLlmTemplatesResponseBodyData {
	s.IsDefaultModel = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetLlmCode(v string) *ListLlmTemplatesResponseBodyData {
	s.LlmCode = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetLlmTemplateId(v string) *ListLlmTemplatesResponseBodyData {
	s.LlmTemplateId = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetName(v string) *ListLlmTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) SetProviderTemplateId(v string) *ListLlmTemplatesResponseBodyData {
	s.ProviderTemplateId = &v
	return s
}

func (s *ListLlmTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
