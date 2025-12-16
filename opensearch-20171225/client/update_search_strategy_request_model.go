// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSearchStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SearchStrategy) *UpdateSearchStrategyRequest
	GetBody() *SearchStrategy
}

type UpdateSearchStrategyRequest struct {
	// The request body.
	Body *SearchStrategy `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSearchStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSearchStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyRequest) GetBody() *SearchStrategy {
	return s.Body
}

func (s *UpdateSearchStrategyRequest) SetBody(v *SearchStrategy) *UpdateSearchStrategyRequest {
	s.Body = v
	return s
}

func (s *UpdateSearchStrategyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
