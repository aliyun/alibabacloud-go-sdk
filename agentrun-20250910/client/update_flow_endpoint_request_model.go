// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateFlowEndpointInput) *UpdateFlowEndpointRequest
	GetBody() *UpdateFlowEndpointInput
}

type UpdateFlowEndpointRequest struct {
	// 包含要更新的工作流端点配置信息的请求体
	//
	// This parameter is required.
	Body *UpdateFlowEndpointInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowEndpointRequest) GetBody() *UpdateFlowEndpointInput {
	return s.Body
}

func (s *UpdateFlowEndpointRequest) SetBody(v *UpdateFlowEndpointInput) *UpdateFlowEndpointRequest {
	s.Body = v
	return s
}

func (s *UpdateFlowEndpointRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
