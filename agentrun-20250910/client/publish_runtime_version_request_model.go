// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRuntimeVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PublishRuntimeVersionInput) *PublishRuntimeVersionRequest
	GetBody() *PublishRuntimeVersionInput
}

type PublishRuntimeVersionRequest struct {
	Body *PublishRuntimeVersionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishRuntimeVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishRuntimeVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishRuntimeVersionRequest) GetBody() *PublishRuntimeVersionInput {
	return s.Body
}

func (s *PublishRuntimeVersionRequest) SetBody(v *PublishRuntimeVersionInput) *PublishRuntimeVersionRequest {
	s.Body = v
	return s
}

func (s *PublishRuntimeVersionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
