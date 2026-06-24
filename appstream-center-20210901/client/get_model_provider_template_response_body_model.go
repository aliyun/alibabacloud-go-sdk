// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProviderTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetModelProviderTemplateResponseBodyData) *GetModelProviderTemplateResponseBody
	GetData() *GetModelProviderTemplateResponseBodyData
	SetRequestId(v string) *GetModelProviderTemplateResponseBody
	GetRequestId() *string
}

type GetModelProviderTemplateResponseBody struct {
	// The returned data object.
	Data *GetModelProviderTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModelProviderTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelProviderTemplateResponseBody) GetData() *GetModelProviderTemplateResponseBodyData {
	return s.Data
}

func (s *GetModelProviderTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelProviderTemplateResponseBody) SetData(v *GetModelProviderTemplateResponseBodyData) *GetModelProviderTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetModelProviderTemplateResponseBody) SetRequestId(v string) *GetModelProviderTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelProviderTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetModelProviderTemplateResponseBodyData struct {
	// The JSON configuration object of the model provider template.
	//
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
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the model provider template.
	//
	// example:
	//
	// 阿里云百炼服务商。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the WUYING secure gateway proxy is enabled.
	//
	// example:
	//
	// true
	EnableWuyingProxy *bool `json:"EnableWuyingProxy,omitempty" xml:"EnableWuyingProxy,omitempty"`
	// The name of the model provider template.
	//
	// example:
	//
	// 阿里云百炼。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the model provider.
	//
	// example:
	//
	// bailian
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The ID of the model provider template.
	//
	// example:
	//
	// mpt-xxxx
	ProviderTemplateId *string `json:"ProviderTemplateId,omitempty" xml:"ProviderTemplateId,omitempty"`
}

func (s GetModelProviderTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelProviderTemplateResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *GetModelProviderTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetModelProviderTemplateResponseBodyData) GetEnableWuyingProxy() *bool {
	return s.EnableWuyingProxy
}

func (s *GetModelProviderTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetModelProviderTemplateResponseBodyData) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetModelProviderTemplateResponseBodyData) GetProviderTemplateId() *string {
	return s.ProviderTemplateId
}

func (s *GetModelProviderTemplateResponseBodyData) SetConfig(v string) *GetModelProviderTemplateResponseBodyData {
	s.Config = &v
	return s
}

func (s *GetModelProviderTemplateResponseBodyData) SetDescription(v string) *GetModelProviderTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetModelProviderTemplateResponseBodyData) SetEnableWuyingProxy(v bool) *GetModelProviderTemplateResponseBodyData {
	s.EnableWuyingProxy = &v
	return s
}

func (s *GetModelProviderTemplateResponseBodyData) SetName(v string) *GetModelProviderTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetModelProviderTemplateResponseBodyData) SetProviderName(v string) *GetModelProviderTemplateResponseBodyData {
	s.ProviderName = &v
	return s
}

func (s *GetModelProviderTemplateResponseBodyData) SetProviderTemplateId(v string) *GetModelProviderTemplateResponseBodyData {
	s.ProviderTemplateId = &v
	return s
}

func (s *GetModelProviderTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
