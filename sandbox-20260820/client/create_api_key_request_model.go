// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateApiKeyInput) *CreateApiKeyRequest
	GetBody() *CreateApiKeyInput
}

type CreateApiKeyRequest struct {
	Body *CreateApiKeyInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequest) GetBody() *CreateApiKeyInput {
	return s.Body
}

func (s *CreateApiKeyRequest) SetBody(v *CreateApiKeyInput) *CreateApiKeyRequest {
	s.Body = v
	return s
}

func (s *CreateApiKeyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
