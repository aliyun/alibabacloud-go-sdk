// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotStreamResponseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateCopilotStreamResponseResponseBody
	GetCode() *string
	SetData(v string) *GenerateCopilotStreamResponseResponseBody
	GetData() *string
	SetMessage(v string) *GenerateCopilotStreamResponseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateCopilotStreamResponseResponseBody
	GetRequestId() *string
}

type GenerateCopilotStreamResponseResponseBody struct {
	// The error code.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned by the LLM service. The data is of the string type. If it is a dict, convert and parse it yourself.
	//
	// example:
	//
	// {\\"choices\\": [{\\"finish_reason\\": \\"stop\\", \\"index\\": 0, \\"message\\": {\\"content\\": \\"Alinux is a Linux-based open-source operating system primarily developed and maintained by Alibaba Group, designed specifically for large-scale cloud computing and big data processing environments. The system has been deeply customized and optimized for high concurrency, low latency, and resource utilization to meet the complex and massive business demands within Alibaba. Due to its excellent performance and stability, Alinux has gradually been adopted by external enterprises and research institutions, particularly in areas that require handling high-load scenarios. Compared to standard Linux distributions, Alinux may include specific kernel patches, optimization toolsets, and other features developed internally by Alibaba. However, it is worth noting that although the "Ali" in the name may suggest a close relationship with Alibaba, Alinux as a project became a sub-project under the OpenAtom Foundation in 2021, known as OpenAnolis, marking its progression toward community-driven and broader adoption.\\", \\"role\\": \\"assistant\\"}}], \\"model\\": \\"Qwen1.5\\", \\"sentiment\\": [], \\"status_code\\": 200, \\"time\\": 6.836989402770996, \\"usage\\": {\\"completion_tokens\\": 180, \\"prompt_tokens\\": 176, \\"time_dict\\": {\\"auto_regression_time\\": 0.03798324399524265, \\"speed\\": 26.327398474054736, \\"total_generate_time\\": 6.836983919143677}, \\"total_tokens\\": 356}}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code description. This field is empty if no error occurs.
	//
	// example:
	//
	// Requests for llm service failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotStreamResponseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotStreamResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateCopilotStreamResponseResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateCopilotStreamResponseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateCopilotStreamResponseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCopilotStreamResponseResponseBody) SetCode(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetData(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetMessage(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetRequestId(v string) *GenerateCopilotStreamResponseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) Validate() error {
	return dara.Validate(s)
}
