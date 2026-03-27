// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUmodelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateUmodelRequest
	GetDescription() *string
}

type UpdateUmodelRequest struct {
	// Description.
	//
	// example:
	//
	// workspace test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateUmodelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUmodelRequest) GoString() string {
	return s.String()
}

func (s *UpdateUmodelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateUmodelRequest) SetDescription(v string) *UpdateUmodelRequest {
	s.Description = &v
	return s
}

func (s *UpdateUmodelRequest) Validate() error {
	return dara.Validate(s)
}
