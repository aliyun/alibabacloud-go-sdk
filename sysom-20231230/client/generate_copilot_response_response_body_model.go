// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCopilotResponseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateCopilotResponseResponseBody
	GetCode() *string
	SetData(v string) *GenerateCopilotResponseResponseBody
	GetData() *string
	SetMassage(v string) *GenerateCopilotResponseResponseBody
	GetMassage() *string
	SetRequestId(v string) *GenerateCopilotResponseResponseBody
	GetRequestId() *string
}

type GenerateCopilotResponseResponseBody struct {
	// The error code.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned by the LLM service. The value is of the string type. If the value is a dict, convert and parse it on your own.
	//
	// example:
	//
	// {\\"choices\\": [{\\"finish_reason\\": \\"stop\\", \\"index\\": 0, \\"message\\": {\\"content\\": \\"Alinux is an open-source operating system based on Linux. It is primarily developed and maintained by Alibaba Group and is designed for large-scale cloud computing and big data processing environments. The system has been deeply customized and optimized for high concurrency, low latency, and resource utilization to meet the complex and massive business requirements within Alibaba. Due to its excellent performance and stability, Alinux has gradually been adopted by external enterprises and research institutions, especially in scenarios that require handling high workloads. Compared with standard Linux distributions, Alinux may include specific kernel patches, optimization toolsets, and other features developed internally by Alibaba. However, it is worth noting that although the "Ali" in the name may suggest a close relationship with Alibaba, Alinux as a project became a sub-project under the OpenAtom Foundation in 2021, known as OpenAnolis, marking its progress toward community-driven and broader adoption.\\", \\"role\\": \\"assistant\\"}}], \\"model\\": \\"Qwen1.5\\", \\"sentiment\\": [], \\"status_code\\": 200, \\"time\\": 6.836989402770996, \\"usage\\": {\\"completion_tokens\\": 180, \\"prompt_tokens\\": 176, \\"time_dict\\": {\\"auto_regression_time\\": 0.03798324399524265, \\"speed\\": 26.327398474054736, \\"total_generate_time\\": 6.836983919143677}, \\"total_tokens\\": 356}}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The error code description. This value is empty if no error occurs.
	//
	// example:
	//
	// Requests for llm service failed
	Massage *string `json:"massage,omitempty" xml:"massage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotResponseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCopilotResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateCopilotResponseResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateCopilotResponseResponseBody) GetMassage() *string {
	return s.Massage
}

func (s *GenerateCopilotResponseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCopilotResponseResponseBody) SetCode(v string) *GenerateCopilotResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetData(v string) *GenerateCopilotResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetMassage(v string) *GenerateCopilotResponseResponseBody {
	s.Massage = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetRequestId(v string) *GenerateCopilotResponseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) Validate() error {
	return dara.Validate(s)
}
