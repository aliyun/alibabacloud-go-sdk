// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateMemoryCollectionInput) *UpdateMemoryCollectionRequest
	GetBody() *UpdateMemoryCollectionInput
}

type UpdateMemoryCollectionRequest struct {
	Body *UpdateMemoryCollectionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryCollectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryCollectionRequest) GetBody() *UpdateMemoryCollectionInput {
	return s.Body
}

func (s *UpdateMemoryCollectionRequest) SetBody(v *UpdateMemoryCollectionInput) *UpdateMemoryCollectionRequest {
	s.Body = v
	return s
}

func (s *UpdateMemoryCollectionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
