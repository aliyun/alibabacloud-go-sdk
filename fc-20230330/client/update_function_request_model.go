// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateFunctionInput) *UpdateFunctionRequest
	GetBody() *UpdateFunctionInput
}

type UpdateFunctionRequest struct {
	// The function information
	//
	// This parameter is required.
	Body *UpdateFunctionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) GetBody() *UpdateFunctionInput {
	return s.Body
}

func (s *UpdateFunctionRequest) SetBody(v *UpdateFunctionInput) *UpdateFunctionRequest {
	s.Body = v
	return s
}

func (s *UpdateFunctionRequest) Validate() error {
	return dara.Validate(s)
}
