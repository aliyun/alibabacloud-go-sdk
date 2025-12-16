// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SearchStrategy) *CreateSearchStrategyRequest
	GetBody() *SearchStrategy
}

type CreateSearchStrategyRequest struct {
	// The query policy.
	Body *SearchStrategy `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyRequest) GetBody() *SearchStrategy {
	return s.Body
}

func (s *CreateSearchStrategyRequest) SetBody(v *SearchStrategy) *CreateSearchStrategyRequest {
	s.Body = v
	return s
}

func (s *CreateSearchStrategyRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
