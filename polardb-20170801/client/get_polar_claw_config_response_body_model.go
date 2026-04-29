// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarClawConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetPolarClawConfigResponseBody
	GetApplicationId() *string
	SetCode(v int32) *GetPolarClawConfigResponseBody
	GetCode() *int32
	SetConfig(v map[string]interface{}) *GetPolarClawConfigResponseBody
	GetConfig() map[string]interface{}
	SetHash(v string) *GetPolarClawConfigResponseBody
	GetHash() *string
	SetMessage(v string) *GetPolarClawConfigResponseBody
	GetMessage() *string
	SetOpenclawVersion(v string) *GetPolarClawConfigResponseBody
	GetOpenclawVersion() *string
	SetRequestId(v string) *GetPolarClawConfigResponseBody
	GetRequestId() *string
}

type GetPolarClawConfigResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//     "models": "{\\"mode\\":\\"merge\\",\\"providers\\":{\\"bailian\\":{\\"models\\":[{\\"input\\":[\\"text\\",\\"image\\"],\\"name\\":\\"qwen3.6-plus\\",\\"api\\":\\"openai-completions\\",\\"id\\":\\"qwen3.6-plus\\",\\"compat\\":{\\"supportsUsageInStreaming\\":true}}],\\"baseUrl\\":\\"https://dashscope.aliyuncs.com/compatible-mode/v1\\",\\"apiKey\\":\\"__OPENCLAW_REDACTED__\\",\\"api\\":\\"openai-completions\\"},\\"polardbCustom\\":{\\"models\\":[{\\"name\\":\\"qwen3-max\\",\\"api\\":\\"openai-completions\\",\\"id\\":\\"qwen3-max\\"}],\\"baseUrl\\":\\"https://dashscope.aliyuncs.com/compatible-mode/v1\\",\\"apiKey\\":\\"__OPENCLAW_REDACTED__\\",\\"api\\":\\"openai-completions\\"},\\"polardb\\":{\\"models\\":[],\\"baseUrl\\":\\"https://dashscope.aliyuncs.com/compatible-mode/v1\\",\\"api\\":\\"openai-completions\\"}}}"
	//
	// }
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 005b55a8e870aaf866598e48a6af0bdbfa3fef704770c5e1cbad99648beaa661
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2026.3.28
	OpenclawVersion *string `json:"OpenclawVersion,omitempty" xml:"OpenclawVersion,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolarClawConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPolarClawConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolarClawConfigResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetPolarClawConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPolarClawConfigResponseBody) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *GetPolarClawConfigResponseBody) GetHash() *string {
	return s.Hash
}

func (s *GetPolarClawConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPolarClawConfigResponseBody) GetOpenclawVersion() *string {
	return s.OpenclawVersion
}

func (s *GetPolarClawConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPolarClawConfigResponseBody) SetApplicationId(v string) *GetPolarClawConfigResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *GetPolarClawConfigResponseBody) SetCode(v int32) *GetPolarClawConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetPolarClawConfigResponseBody) SetConfig(v map[string]interface{}) *GetPolarClawConfigResponseBody {
	s.Config = v
	return s
}

func (s *GetPolarClawConfigResponseBody) SetHash(v string) *GetPolarClawConfigResponseBody {
	s.Hash = &v
	return s
}

func (s *GetPolarClawConfigResponseBody) SetMessage(v string) *GetPolarClawConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetPolarClawConfigResponseBody) SetOpenclawVersion(v string) *GetPolarClawConfigResponseBody {
	s.OpenclawVersion = &v
	return s
}

func (s *GetPolarClawConfigResponseBody) SetRequestId(v string) *GetPolarClawConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolarClawConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
