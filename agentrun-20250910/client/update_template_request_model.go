// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateTemplateInput) *UpdateTemplateRequest
	GetBody() *UpdateTemplateInput
	SetClientToken(v string) *UpdateTemplateRequest
	GetClientToken() *string
}

type UpdateTemplateRequest struct {
	// 更新模板所需的配置信息
	//
	// This parameter is required.
	Body *UpdateTemplateInput `json:"body,omitempty" xml:"body,omitempty"`
	// 用于确保请求幂等性的唯一标识符
	//
	// example:
	//
	// clientToken
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetBody() *UpdateTemplateInput {
	return s.Body
}

func (s *UpdateTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTemplateRequest) SetBody(v *UpdateTemplateInput) *UpdateTemplateRequest {
	s.Body = v
	return s
}

func (s *UpdateTemplateRequest) SetClientToken(v string) *UpdateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
