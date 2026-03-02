// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeveloperRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Developer) *CreateDeveloperRequest
	GetBody() *Developer
}

type CreateDeveloperRequest struct {
	// This parameter is required.
	Body *Developer `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeveloperRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeveloperRequest) GoString() string {
	return s.String()
}

func (s *CreateDeveloperRequest) GetBody() *Developer {
	return s.Body
}

func (s *CreateDeveloperRequest) SetBody(v *Developer) *CreateDeveloperRequest {
	s.Body = v
	return s
}

func (s *CreateDeveloperRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
