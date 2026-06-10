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
	// error code
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Data returned by the LLM service, in string format. If it represents a dictionary, you must convert and parse it yourself.
	//
	// example:
	//
	// {\\"choices\\": [{\\"finish_reason\\": \\"stop\\", \\"index\\": 0, \\"message\\": {\\"content\\": \\"Alinux是一个基于Linux的开源操作系统，它主要由阿里巴巴集团开发和维护，专为大规模云计算和大数据处理环境设计。这个系统针对高并发、低延迟以及资源利用率优化进行了深度定制和改良，以适应阿里内部复杂且庞大的业务需求。由于其在性能和稳定性方面的优秀表现，Alinux也在逐渐被外部的一些企业和研究机构采用，特别是在需要处理高负载场景的领域。与标准Linux发行版相比，Alinux可能包含特定的内核补丁、优化工具集以及其他阿里巴巴内部研发的特性。然而，值得注意的是，尽管名称中的“Ali”可能让人联想到与阿里巴巴的紧密关系，但Alinux作为一个项目，已经于2021年成为开放原子开源基金会（OpenAtom Foundation）下的一个子项目，即 OpenAnolis，这标志着它的社区化和广泛化的进程。\\", \\"role\\": \\"assistant\\"}}], \\"model\\": \\"Qwen1.5\\", \\"sentiment\\": [], \\"status_code\\": 200, \\"time\\": 6.836989402770996, \\"usage\\": {\\"completion_tokens\\": 180, \\"prompt_tokens\\": 176, \\"time_dict\\": {\\"auto_regression_time\\": 0.03798324399524265, \\"speed\\": 26.327398474054736, \\"total_generate_time\\": 6.836983919143677}, \\"total_tokens\\": 356}}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Description of the error code. This field is empty if no error occurs.
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
