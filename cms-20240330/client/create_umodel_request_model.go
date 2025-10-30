// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateUmodelRequest
	GetDescription() *string
}

type CreateUmodelRequest struct {
	// Umodel description
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUmodelRequest) GoString() string {
	return s.String()
}

func (s *CreateUmodelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUmodelRequest) SetDescription(v string) *CreateUmodelRequest {
	s.Description = &v
	return s
}

func (s *CreateUmodelRequest) Validate() error {
	return dara.Validate(s)
}
