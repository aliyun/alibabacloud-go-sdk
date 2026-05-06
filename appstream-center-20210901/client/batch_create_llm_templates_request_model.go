// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateLlmTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLlmTemplateItems(v []*BatchCreateLlmTemplatesRequestLlmTemplateItems) *BatchCreateLlmTemplatesRequest
	GetLlmTemplateItems() []*BatchCreateLlmTemplatesRequestLlmTemplateItems
	SetModelTemplateId(v string) *BatchCreateLlmTemplatesRequest
	GetModelTemplateId() *string
	SetProviderTemplateId(v string) *BatchCreateLlmTemplatesRequest
	GetProviderTemplateId() *string
}

type BatchCreateLlmTemplatesRequest struct {
	LlmTemplateItems []*BatchCreateLlmTemplatesRequestLlmTemplateItems `json:"LlmTemplateItems,omitempty" xml:"LlmTemplateItems,omitempty" type:"Repeated"`
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s BatchCreateLlmTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateLlmTemplatesRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateLlmTemplatesRequest) GetLlmTemplateItems() []*BatchCreateLlmTemplatesRequestLlmTemplateItems {
	return s.LlmTemplateItems
}

func (s *BatchCreateLlmTemplatesRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *BatchCreateLlmTemplatesRequest) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *BatchCreateLlmTemplatesRequest) SetLlmTemplateItems(v []*BatchCreateLlmTemplatesRequestLlmTemplateItems) *BatchCreateLlmTemplatesRequest {
	s.LlmTemplateItems = v
	return s
}

func (s *BatchCreateLlmTemplatesRequest) SetModelTemplateId(v string) *BatchCreateLlmTemplatesRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequest) SetProviderTemplateId(v string) *BatchCreateLlmTemplatesRequest {
	s.ProviderTemplateId = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequest) Validate() error {
	if s.LlmTemplateItems != nil {
		for _, item := range s.LlmTemplateItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateLlmTemplatesRequestLlmTemplateItems struct {
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
	Config      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	IsDefaultModel *bool `json:"IsDefaultModel,omitempty" xml:"IsDefaultModel,omitempty"`
	// example:
	//
	// qwen3.6-plus
	LlmCode *string `json:"LlmCode,omitempty" xml:"LlmCode,omitempty"`
	// example:
	//
	// Qwen3.6-Plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchCreateLlmTemplatesRequestLlmTemplateItems) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateLlmTemplatesRequestLlmTemplateItems) GoString() string {
	return s.String()
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) GetConfig() *string {
	return s.Config
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) GetDescription() *string {
	return s.Description
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) GetIsDefaultModel() *bool {
	return s.IsDefaultModel
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) GetLlmCode() *string {
	return s.LlmCode
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) GetName() *string {
	return s.Name
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) SetConfig(v string) *BatchCreateLlmTemplatesRequestLlmTemplateItems {
	s.Config = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) SetDescription(v string) *BatchCreateLlmTemplatesRequestLlmTemplateItems {
	s.Description = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) SetIsDefaultModel(v bool) *BatchCreateLlmTemplatesRequestLlmTemplateItems {
	s.IsDefaultModel = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) SetLlmCode(v string) *BatchCreateLlmTemplatesRequestLlmTemplateItems {
	s.LlmCode = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) SetName(v string) *BatchCreateLlmTemplatesRequestLlmTemplateItems {
	s.Name = &v
	return s
}

func (s *BatchCreateLlmTemplatesRequestLlmTemplateItems) Validate() error {
	return dara.Validate(s)
}
