// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateFlowInput) *CreateFlowRequest
	GetBody() *CreateFlowInput
}

type CreateFlowRequest struct {
	// 创建工作流所需的完整配置信息，包括工作流名称、定义、执行模式等
	//
	// This parameter is required.
	Body *CreateFlowInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) GetBody() *CreateFlowInput {
	return s.Body
}

func (s *CreateFlowRequest) SetBody(v *CreateFlowInput) *CreateFlowRequest {
	s.Body = v
	return s
}

func (s *CreateFlowRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
