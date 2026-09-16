// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDigitalEmployeeUmodelRequest
	GetDescription() *string
}

type UpdateDigitalEmployeeUmodelRequest struct {
	// The updated UModel description of the digital human.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateDigitalEmployeeUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeUmodelRequest) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeUmodelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDigitalEmployeeUmodelRequest) SetDescription(v string) *UpdateDigitalEmployeeUmodelRequest {
	s.Description = &v
	return s
}

func (s *UpdateDigitalEmployeeUmodelRequest) Validate() error {
	return dara.Validate(s)
}
