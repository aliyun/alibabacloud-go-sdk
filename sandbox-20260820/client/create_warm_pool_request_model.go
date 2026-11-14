// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarmPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateWarmPoolInput) *CreateWarmPoolRequest
	GetBody() *CreateWarmPoolInput
}

type CreateWarmPoolRequest struct {
	// The request body, which contains the warm pool configuration.
	Body *CreateWarmPoolInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarmPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWarmPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateWarmPoolRequest) GetBody() *CreateWarmPoolInput {
	return s.Body
}

func (s *CreateWarmPoolRequest) SetBody(v *CreateWarmPoolInput) *CreateWarmPoolRequest {
	s.Body = v
	return s
}

func (s *CreateWarmPoolRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
