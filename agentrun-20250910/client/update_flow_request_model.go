// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateFlowInput) *UpdateFlowRequest
	GetBody() *UpdateFlowInput
}

type UpdateFlowRequest struct {
	// 包含要更新的工作流配置信息的请求体
	//
	// This parameter is required.
	Body *UpdateFlowInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequest) GetBody() *UpdateFlowInput {
	return s.Body
}

func (s *UpdateFlowRequest) SetBody(v *UpdateFlowInput) *UpdateFlowRequest {
	s.Body = v
	return s
}

func (s *UpdateFlowRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
