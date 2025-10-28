// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutConcurrencyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PutConcurrencyInput) *PutConcurrencyConfigRequest
	GetBody() *PutConcurrencyInput
}

type PutConcurrencyConfigRequest struct {
	// The concurrency configurations.
	//
	// This parameter is required.
	Body *PutConcurrencyInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutConcurrencyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PutConcurrencyConfigRequest) GoString() string {
	return s.String()
}

func (s *PutConcurrencyConfigRequest) GetBody() *PutConcurrencyInput {
	return s.Body
}

func (s *PutConcurrencyConfigRequest) SetBody(v *PutConcurrencyInput) *PutConcurrencyConfigRequest {
	s.Body = v
	return s
}

func (s *PutConcurrencyConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
