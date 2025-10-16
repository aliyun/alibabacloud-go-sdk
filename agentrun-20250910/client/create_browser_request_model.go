// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBrowserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateBrowserInput) *CreateBrowserRequest
	GetBody() *CreateBrowserInput
}

type CreateBrowserRequest struct {
	// This parameter is required.
	Body *CreateBrowserInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBrowserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBrowserRequest) GoString() string {
	return s.String()
}

func (s *CreateBrowserRequest) GetBody() *CreateBrowserInput {
	return s.Body
}

func (s *CreateBrowserRequest) SetBody(v *CreateBrowserInput) *CreateBrowserRequest {
	s.Body = v
	return s
}

func (s *CreateBrowserRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
