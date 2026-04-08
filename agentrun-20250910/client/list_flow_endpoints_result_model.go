// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowEndpointsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlowEndpointsResult
	GetCode() *string
	SetData(v *ListFlowEndpointsOutput) *ListFlowEndpointsResult
	GetData() *ListFlowEndpointsOutput
	SetRequestId(v string) *ListFlowEndpointsResult
	GetRequestId() *string
}

type ListFlowEndpointsResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工作流端点列表的详细信息
	//
	// example:
	//
	// {}
	Data *ListFlowEndpointsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFlowEndpointsResult) String() string {
	return dara.Prettify(s)
}

func (s ListFlowEndpointsResult) GoString() string {
	return s.String()
}

func (s *ListFlowEndpointsResult) GetCode() *string {
	return s.Code
}

func (s *ListFlowEndpointsResult) GetData() *ListFlowEndpointsOutput {
	return s.Data
}

func (s *ListFlowEndpointsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlowEndpointsResult) SetCode(v string) *ListFlowEndpointsResult {
	s.Code = &v
	return s
}

func (s *ListFlowEndpointsResult) SetData(v *ListFlowEndpointsOutput) *ListFlowEndpointsResult {
	s.Data = v
	return s
}

func (s *ListFlowEndpointsResult) SetRequestId(v string) *ListFlowEndpointsResult {
	s.RequestId = &v
	return s
}

func (s *ListFlowEndpointsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
