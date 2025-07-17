// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperations(v []*HttpApiOperation) *CreateHttpApiOperationRequest
	GetOperations() []*HttpApiOperation
}

type CreateHttpApiOperationRequest struct {
	// List of operation definitions.
	Operations []*HttpApiOperation `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
}

func (s CreateHttpApiOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiOperationRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationRequest) GetOperations() []*HttpApiOperation {
	return s.Operations
}

func (s *CreateHttpApiOperationRequest) SetOperations(v []*HttpApiOperation) *CreateHttpApiOperationRequest {
	s.Operations = v
	return s
}

func (s *CreateHttpApiOperationRequest) Validate() error {
	return dara.Validate(s)
}
