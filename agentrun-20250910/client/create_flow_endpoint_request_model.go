// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateFlowEndpointInput) *CreateFlowEndpointRequest
	GetBody() *CreateFlowEndpointInput
}

type CreateFlowEndpointRequest struct {
	// 创建工作流端点所需的完整配置信息
	//
	// This parameter is required.
	Body *CreateFlowEndpointInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowEndpointRequest) GetBody() *CreateFlowEndpointInput {
	return s.Body
}

func (s *CreateFlowEndpointRequest) SetBody(v *CreateFlowEndpointInput) *CreateFlowEndpointRequest {
	s.Body = v
	return s
}

func (s *CreateFlowEndpointRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
