// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLLMConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLLMConfigId(v string) *CreateLLMConfigResponseBody
	GetLLMConfigId() *string
	SetRequestId(v string) *CreateLLMConfigResponseBody
	GetRequestId() *string
}

type CreateLLMConfigResponseBody struct {
	// example:
	//
	// llm_config1
	LLMConfigId *string `json:"LLMConfigId,omitempty" xml:"LLMConfigId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLLMConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLLMConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLLMConfigResponseBody) GetLLMConfigId() *string {
	return s.LLMConfigId
}

func (s *CreateLLMConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLLMConfigResponseBody) SetLLMConfigId(v string) *CreateLLMConfigResponseBody {
	s.LLMConfigId = &v
	return s
}

func (s *CreateLLMConfigResponseBody) SetRequestId(v string) *CreateLLMConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLLMConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
