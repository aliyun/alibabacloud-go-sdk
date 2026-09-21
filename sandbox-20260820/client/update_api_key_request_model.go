// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateApiKeyInput) *UpdateApiKeyRequest
	GetBody() *UpdateApiKeyInput
}

type UpdateApiKeyRequest struct {
	Body *UpdateApiKeyInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyRequest) GetBody() *UpdateApiKeyInput {
	return s.Body
}

func (s *UpdateApiKeyRequest) SetBody(v *UpdateApiKeyInput) *UpdateApiKeyRequest {
	s.Body = v
	return s
}

func (s *UpdateApiKeyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
