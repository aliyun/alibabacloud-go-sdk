// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableLLMModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModels(v []*string) *GetAvailableLLMModelsResponseBody
	GetModels() []*string
	SetRequestId(v string) *GetAvailableLLMModelsResponseBody
	GetRequestId() *string
}

type GetAvailableLLMModelsResponseBody struct {
	// example:
	//
	// ["glm-5","deepseek-v3.2"]
	Models []*string `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAvailableLLMModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableLLMModelsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAvailableLLMModelsResponseBody) GetModels() []*string {
	return s.Models
}

func (s *GetAvailableLLMModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAvailableLLMModelsResponseBody) SetModels(v []*string) *GetAvailableLLMModelsResponseBody {
	s.Models = v
	return s
}

func (s *GetAvailableLLMModelsResponseBody) SetRequestId(v string) *GetAvailableLLMModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAvailableLLMModelsResponseBody) Validate() error {
	return dara.Validate(s)
}
