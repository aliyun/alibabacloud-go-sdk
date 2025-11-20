// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateModelProxyInput) *UpdateModelProxyRequest
	GetBody() *UpdateModelProxyInput
}

type UpdateModelProxyRequest struct {
	Body *UpdateModelProxyInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelProxyRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelProxyRequest) GetBody() *UpdateModelProxyInput {
	return s.Body
}

func (s *UpdateModelProxyRequest) SetBody(v *UpdateModelProxyInput) *UpdateModelProxyRequest {
	s.Body = v
	return s
}

func (s *UpdateModelProxyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
