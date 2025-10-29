// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v *HttpApiOperation) *UpdateHttpApiOperationRequest
	GetOperation() *HttpApiOperation
}

type UpdateHttpApiOperationRequest struct {
	// operation definition.
	Operation *HttpApiOperation `json:"operation,omitempty" xml:"operation,omitempty"`
}

func (s UpdateHttpApiOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiOperationRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiOperationRequest) GetOperation() *HttpApiOperation {
	return s.Operation
}

func (s *UpdateHttpApiOperationRequest) SetOperation(v *HttpApiOperation) *UpdateHttpApiOperationRequest {
	s.Operation = v
	return s
}

func (s *UpdateHttpApiOperationRequest) Validate() error {
	if s.Operation != nil {
		if err := s.Operation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
