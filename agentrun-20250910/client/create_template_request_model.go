// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateTemplateInput) *CreateTemplateRequest
	GetBody() *CreateTemplateInput
}

type CreateTemplateRequest struct {
	// 创建模板所需的完整配置信息，包括模板名称、类型、资源规格、网络配置等
	//
	// This parameter is required.
	Body *CreateTemplateInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetBody() *CreateTemplateInput {
	return s.Body
}

func (s *CreateTemplateRequest) SetBody(v *CreateTemplateInput) *CreateTemplateRequest {
	s.Body = v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
