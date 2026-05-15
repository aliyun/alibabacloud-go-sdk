// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentResourcesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentResourcesResult
	GetCode() *string
	SetData(v *ListAgentResourcesOutput) *ListAgentResourcesResult
	GetData() *ListAgentResourcesOutput
	SetRequestId(v string) *ListAgentResourcesResult
	GetRequestId() *string
}

type ListAgentResourcesResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 智能体资源列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListAgentResourcesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentResourcesResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentResourcesResult) GoString() string {
	return s.String()
}

func (s *ListAgentResourcesResult) GetCode() *string {
	return s.Code
}

func (s *ListAgentResourcesResult) GetData() *ListAgentResourcesOutput {
	return s.Data
}

func (s *ListAgentResourcesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentResourcesResult) SetCode(v string) *ListAgentResourcesResult {
	s.Code = &v
	return s
}

func (s *ListAgentResourcesResult) SetData(v *ListAgentResourcesOutput) *ListAgentResourcesResult {
	s.Data = v
	return s
}

func (s *ListAgentResourcesResult) SetRequestId(v string) *ListAgentResourcesResult {
	s.RequestId = &v
	return s
}

func (s *ListAgentResourcesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
