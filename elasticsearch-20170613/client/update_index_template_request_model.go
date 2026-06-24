// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIndexTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateIndexTemplateRequest
	GetClientToken() *string
	SetBody(v string) *UpdateIndexTemplateRequest
	GetBody() *string
}

type UpdateIndexTemplateRequest struct {
	// A unique token used to ensure idempotence of the request. The client generates this value. The value must be unique among different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIndexTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIndexTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateIndexTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIndexTemplateRequest) GetBody() *string {
	return s.Body
}

func (s *UpdateIndexTemplateRequest) SetClientToken(v string) *UpdateIndexTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIndexTemplateRequest) SetBody(v string) *UpdateIndexTemplateRequest {
	s.Body = &v
	return s
}

func (s *UpdateIndexTemplateRequest) Validate() error {
	return dara.Validate(s)
}
