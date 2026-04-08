// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowEndpointResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlowEndpointResult
	GetCode() *string
	SetData(v *FlowEndpoint) *FlowEndpointResult
	GetData() *FlowEndpoint
	SetRequestId(v string) *FlowEndpointResult
	GetRequestId() *string
}

type FlowEndpointResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工作流端点的详细信息
	//
	// example:
	//
	// {}
	Data *FlowEndpoint `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FlowEndpointResult) String() string {
	return dara.Prettify(s)
}

func (s FlowEndpointResult) GoString() string {
	return s.String()
}

func (s *FlowEndpointResult) GetCode() *string {
	return s.Code
}

func (s *FlowEndpointResult) GetData() *FlowEndpoint {
	return s.Data
}

func (s *FlowEndpointResult) GetRequestId() *string {
	return s.RequestId
}

func (s *FlowEndpointResult) SetCode(v string) *FlowEndpointResult {
	s.Code = &v
	return s
}

func (s *FlowEndpointResult) SetData(v *FlowEndpoint) *FlowEndpointResult {
	s.Data = v
	return s
}

func (s *FlowEndpointResult) SetRequestId(v string) *FlowEndpointResult {
	s.RequestId = &v
	return s
}

func (s *FlowEndpointResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
