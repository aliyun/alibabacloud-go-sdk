// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarmPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateWarmPoolInput) *UpdateWarmPoolRequest
	GetBody() *UpdateWarmPoolInput
}

type UpdateWarmPoolRequest struct {
	// The request body, which contains the warm pool configuration.
	Body *UpdateWarmPoolInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarmPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarmPoolRequest) GoString() string {
	return s.String()
}

func (s *UpdateWarmPoolRequest) GetBody() *UpdateWarmPoolInput {
	return s.Body
}

func (s *UpdateWarmPoolRequest) SetBody(v *UpdateWarmPoolInput) *UpdateWarmPoolRequest {
	s.Body = v
	return s
}

func (s *UpdateWarmPoolRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
