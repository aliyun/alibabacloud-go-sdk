// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateMemoryCollectionInput) *CreateMemoryCollectionRequest
	GetBody() *CreateMemoryCollectionInput
}

type CreateMemoryCollectionRequest struct {
	Body *CreateMemoryCollectionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryCollectionRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryCollectionRequest) GetBody() *CreateMemoryCollectionInput {
	return s.Body
}

func (s *CreateMemoryCollectionRequest) SetBody(v *CreateMemoryCollectionInput) *CreateMemoryCollectionRequest {
	s.Body = v
	return s
}

func (s *CreateMemoryCollectionRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
