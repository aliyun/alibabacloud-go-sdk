// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowVersionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlowVersionResult
	GetCode() *string
	SetData(v *FlowVersion) *FlowVersionResult
	GetData() *FlowVersion
	SetRequestId(v string) *FlowVersionResult
	GetRequestId() *string
}

type FlowVersionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工作流版本的详细信息
	//
	// example:
	//
	// {}
	Data *FlowVersion `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FlowVersionResult) String() string {
	return dara.Prettify(s)
}

func (s FlowVersionResult) GoString() string {
	return s.String()
}

func (s *FlowVersionResult) GetCode() *string {
	return s.Code
}

func (s *FlowVersionResult) GetData() *FlowVersion {
	return s.Data
}

func (s *FlowVersionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *FlowVersionResult) SetCode(v string) *FlowVersionResult {
	s.Code = &v
	return s
}

func (s *FlowVersionResult) SetData(v *FlowVersion) *FlowVersionResult {
	s.Data = v
	return s
}

func (s *FlowVersionResult) SetRequestId(v string) *FlowVersionResult {
	s.RequestId = &v
	return s
}

func (s *FlowVersionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
