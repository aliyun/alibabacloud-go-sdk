// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlowsResult
	GetCode() *string
	SetData(v *ListFlowsOutput) *ListFlowsResult
	GetData() *ListFlowsOutput
	SetRequestId(v string) *ListFlowsResult
	GetRequestId() *string
}

type ListFlowsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工作流列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListFlowsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFlowsResult) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsResult) GoString() string {
	return s.String()
}

func (s *ListFlowsResult) GetCode() *string {
	return s.Code
}

func (s *ListFlowsResult) GetData() *ListFlowsOutput {
	return s.Data
}

func (s *ListFlowsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowsResult) SetCode(v string) *ListFlowsResult {
	s.Code = &v
	return s
}

func (s *ListFlowsResult) SetData(v *ListFlowsOutput) *ListFlowsResult {
	s.Data = v
	return s
}

func (s *ListFlowsResult) SetRequestId(v string) *ListFlowsResult {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
