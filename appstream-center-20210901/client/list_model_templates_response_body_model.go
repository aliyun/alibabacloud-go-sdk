// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListModelTemplatesResponseBodyData) *ListModelTemplatesResponseBody
	GetData() []*ListModelTemplatesResponseBodyData
	SetPageNumber(v int32) *ListModelTemplatesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplatesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModelTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModelTemplatesResponseBody
	GetTotalCount() *int32
}

type ListModelTemplatesResponseBody struct {
	// The list of returned data objects.
	Data []*ListModelTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the current query result.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in the query result.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the query result.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesResponseBody) GetData() []*ListModelTemplatesResponseBodyData {
	return s.Data
}

func (s *ListModelTemplatesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplatesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelTemplatesResponseBody) SetData(v []*ListModelTemplatesResponseBodyData) *ListModelTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListModelTemplatesResponseBody) SetPageNumber(v int32) *ListModelTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplatesResponseBody) SetPageSize(v int32) *ListModelTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplatesResponseBody) SetRequestId(v string) *ListModelTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelTemplatesResponseBody) SetTotalCount(v int32) *ListModelTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelTemplatesResponseBody) Validate() error {
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

type ListModelTemplatesResponseBodyData struct {
	// The name of the Agent provider.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The model group configuration JSON object.
	//
	// example:
	//
	// {
	//
	// 	"defaults": {
	//
	// 		"model": {
	//
	// 			"primary": "bailian/qwen3.5-plus"
	//
	// 		}
	//
	// 	}
	//
	// }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The template group description.
	//
	// example:
	//
	// 测试模型分组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether models have been configured in the group.
	//
	// example:
	//
	// true
	HasModel *bool `json:"HasModel,omitempty" xml:"HasModel,omitempty"`
	// The model group ID.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The template group name.
	//
	// example:
	//
	// model-template-001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListModelTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesResponseBodyData) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelTemplatesResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *ListModelTemplatesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListModelTemplatesResponseBodyData) GetHasModel() *bool {
	return s.HasModel
}

func (s *ListModelTemplatesResponseBodyData) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListModelTemplatesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListModelTemplatesResponseBodyData) SetAgentProvider(v string) *ListModelTemplatesResponseBodyData {
	s.AgentProvider = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetConfig(v string) *ListModelTemplatesResponseBodyData {
	s.Config = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetDescription(v string) *ListModelTemplatesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetHasModel(v bool) *ListModelTemplatesResponseBodyData {
	s.HasModel = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetModelTemplateId(v string) *ListModelTemplatesResponseBodyData {
	s.ModelTemplateId = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) SetName(v string) *ListModelTemplatesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListModelTemplatesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
