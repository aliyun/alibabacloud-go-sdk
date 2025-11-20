// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateModelProxyInput) *CreateModelProxyRequest
	GetBody() *CreateModelProxyInput
}

type CreateModelProxyRequest struct {
	Body *CreateModelProxyInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelProxyRequest) GoString() string {
	return s.String()
}

func (s *CreateModelProxyRequest) GetBody() *CreateModelProxyInput {
	return s.Body
}

func (s *CreateModelProxyRequest) SetBody(v *CreateModelProxyInput) *CreateModelProxyRequest {
	s.Body = v
	return s
}

func (s *CreateModelProxyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
