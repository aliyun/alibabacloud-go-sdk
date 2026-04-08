// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlowVersionsResult
	GetCode() *string
	SetData(v *ListFlowVersionsOutput) *ListFlowVersionsResult
	GetData() *ListFlowVersionsOutput
	SetRequestId(v string) *ListFlowVersionsResult
	GetRequestId() *string
}

type ListFlowVersionsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工作流版本列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListFlowVersionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFlowVersionsResult) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsResult) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsResult) GetCode() *string {
	return s.Code
}

func (s *ListFlowVersionsResult) GetData() *ListFlowVersionsOutput {
	return s.Data
}

func (s *ListFlowVersionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowVersionsResult) SetCode(v string) *ListFlowVersionsResult {
	s.Code = &v
	return s
}

func (s *ListFlowVersionsResult) SetData(v *ListFlowVersionsOutput) *ListFlowVersionsResult {
	s.Data = v
	return s
}

func (s *ListFlowVersionsResult) SetRequestId(v string) *ListFlowVersionsResult {
	s.RequestId = &v
	return s
}

func (s *ListFlowVersionsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
