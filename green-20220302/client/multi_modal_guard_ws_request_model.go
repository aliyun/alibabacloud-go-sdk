// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardWsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelType(v string) *MultiModalGuardWsRequest
	GetModelType() *string
	SetProtocolType(v string) *MultiModalGuardWsRequest
	GetProtocolType() *string
	SetService(v string) *MultiModalGuardWsRequest
	GetService() *string
	SetServiceParameters(v string) *MultiModalGuardWsRequest
	GetServiceParameters() *string
}

type MultiModalGuardWsRequest struct {
	// The model type. Valid values:
	//
	// - LLM
	//
	// example:
	//
	// LLM
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The protocol type. Valid values:
	//
	// - OpenAI
	//
	// - DashScope
	//
	// - Anthropic
	//
	// example:
	//
	// DashScope
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The moderation service category. Valid values:
	//
	// - query_security_check_pro: AI input content security detection (pro edition).
	//
	// - response_security_check_pro: AI-generated content security detection (pro edition).
	//
	// - img_query_security_check: AIGC input image security detection.
	//
	// - img_response_security_check: AIGC output image security detection.
	//
	// - text_img_mix_guard: Multimodal (text + image) hybrid security detection.
	//
	// - file_security_sync_check: AIGC input or output file security detection.
	//
	// - text_file_sec_sync_check: Multimodal (text + file) real-time security detection.
	//
	// example:
	//
	// query_security_check_pro
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameter set required by the moderation service, in JSON string format. The input parameter for text content is content (String), the input parameter for image content is imageUrls (JSONArray), and the input parameter for file content is fileUrls (JSONArray).
	//
	// example:
	//
	// - 文本：
	//
	// {
	//
	//   "content": "test"
	//
	// }
	//
	// - 图片：
	//
	// {
	//
	//   "imageUrls": ["https://example.com/image.png"]
	//
	// }
	//
	// - 文件：
	//
	// {
	//
	//   "fileUrls": ["https://example.com/file.pdf"]
	//
	// }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalGuardWsRequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardWsRequest) GoString() string {
	return s.String()
}

func (s *MultiModalGuardWsRequest) GetModelType() *string {
	return s.ModelType
}

func (s *MultiModalGuardWsRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *MultiModalGuardWsRequest) GetService() *string {
	return s.Service
}

func (s *MultiModalGuardWsRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalGuardWsRequest) SetModelType(v string) *MultiModalGuardWsRequest {
	s.ModelType = &v
	return s
}

func (s *MultiModalGuardWsRequest) SetProtocolType(v string) *MultiModalGuardWsRequest {
	s.ProtocolType = &v
	return s
}

func (s *MultiModalGuardWsRequest) SetService(v string) *MultiModalGuardWsRequest {
	s.Service = &v
	return s
}

func (s *MultiModalGuardWsRequest) SetServiceParameters(v string) *MultiModalGuardWsRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalGuardWsRequest) Validate() error {
	return dara.Validate(s)
}
