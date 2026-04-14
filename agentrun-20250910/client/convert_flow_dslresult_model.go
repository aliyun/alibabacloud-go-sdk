// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertFlowDSLResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConvertFlowDSLResult
	GetCode() *string
	SetData(v *ConvertFlowDSLData) *ConvertFlowDSLResult
	GetData() *ConvertFlowDSLData
	SetRequestId(v string) *ConvertFlowDSLResult
	GetRequestId() *string
}

type ConvertFlowDSLResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 工作流DSL转换的详细结果数据
	Data *ConvertFlowDSLData `json:"data,omitempty" xml:"data,omitempty"`
	// 唯一的请求标识符，用于问题追踪
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ConvertFlowDSLResult) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLResult) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLResult) GetCode() *string {
	return s.Code
}

func (s *ConvertFlowDSLResult) GetData() *ConvertFlowDSLData {
	return s.Data
}

func (s *ConvertFlowDSLResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertFlowDSLResult) SetCode(v string) *ConvertFlowDSLResult {
	s.Code = &v
	return s
}

func (s *ConvertFlowDSLResult) SetData(v *ConvertFlowDSLData) *ConvertFlowDSLResult {
	s.Data = v
	return s
}

func (s *ConvertFlowDSLResult) SetRequestId(v string) *ConvertFlowDSLResult {
	s.RequestId = &v
	return s
}

func (s *ConvertFlowDSLResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
