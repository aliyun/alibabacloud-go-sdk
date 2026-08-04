// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModelRouterBatchCreateModelRequest
	GetApiKey() *string
	SetBaseUrl(v string) *ModelRouterBatchCreateModelRequest
	GetBaseUrl() *string
	SetModels(v []*BatchCreateModelItemDTO) *ModelRouterBatchCreateModelRequest
	GetModels() []*BatchCreateModelItemDTO
	SetSymbol(v string) *ModelRouterBatchCreateModelRequest
	GetSymbol() *string
}

type ModelRouterBatchCreateModelRequest struct {
	// The API key. This parameter is required. The key is shared by the same provider and reused by all models.
	//
	// This parameter is required.
	//
	// example:
	//
	// sk-xxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The base URL. This parameter is optional. Specify this parameter when you use a custom gateway address. If you do not specify this parameter, the default address of the provider is used.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// The list of models to create in batches. This parameter is required. At least one item must be specified.
	//
	// This parameter is required.
	Models []*BatchCreateModelItemDTO `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// The provider symbol. This parameter is required. All models items share the same provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen
	Symbol *string `json:"symbol,omitempty" xml:"symbol,omitempty"`
}

func (s ModelRouterBatchCreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateModelRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateModelRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelRouterBatchCreateModelRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelRouterBatchCreateModelRequest) GetModels() []*BatchCreateModelItemDTO {
	return s.Models
}

func (s *ModelRouterBatchCreateModelRequest) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterBatchCreateModelRequest) SetApiKey(v string) *ModelRouterBatchCreateModelRequest {
	s.ApiKey = &v
	return s
}

func (s *ModelRouterBatchCreateModelRequest) SetBaseUrl(v string) *ModelRouterBatchCreateModelRequest {
	s.BaseUrl = &v
	return s
}

func (s *ModelRouterBatchCreateModelRequest) SetModels(v []*BatchCreateModelItemDTO) *ModelRouterBatchCreateModelRequest {
	s.Models = v
	return s
}

func (s *ModelRouterBatchCreateModelRequest) SetSymbol(v string) *ModelRouterBatchCreateModelRequest {
	s.Symbol = &v
	return s
}

func (s *ModelRouterBatchCreateModelRequest) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
