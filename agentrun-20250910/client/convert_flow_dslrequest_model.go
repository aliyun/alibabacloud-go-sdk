// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertFlowDSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ConvertFlowDSLInput) *ConvertFlowDSLRequest
	GetBody() *ConvertFlowDSLInput
}

type ConvertFlowDSLRequest struct {
	// 包含待转换的工作流DSL内容和转换选项，支持多种DSL格式（如Dify、n8n等）以及inline和base64两种输入方式
	//
	// This parameter is required.
	Body *ConvertFlowDSLInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertFlowDSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLRequest) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLRequest) GetBody() *ConvertFlowDSLInput {
	return s.Body
}

func (s *ConvertFlowDSLRequest) SetBody(v *ConvertFlowDSLInput) *ConvertFlowDSLRequest {
	s.Body = v
	return s
}

func (s *ConvertFlowDSLRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
