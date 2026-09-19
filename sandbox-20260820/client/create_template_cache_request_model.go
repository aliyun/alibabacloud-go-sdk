// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateTemplateCacheInput) *CreateTemplateCacheRequest
	GetBody() *CreateTemplateCacheInput
}

type CreateTemplateCacheRequest struct {
	Body *CreateTemplateCacheInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateCacheRequest) GetBody() *CreateTemplateCacheInput {
	return s.Body
}

func (s *CreateTemplateCacheRequest) SetBody(v *CreateTemplateCacheInput) *CreateTemplateCacheRequest {
	s.Body = v
	return s
}

func (s *CreateTemplateCacheRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
