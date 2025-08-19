// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateFunctionInput) *CreateFunctionRequest
	GetBody() *CreateFunctionInput
}

type CreateFunctionRequest struct {
	// The information about function configurations.
	//
	// This parameter is required.
	Body *CreateFunctionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) GetBody() *CreateFunctionInput {
	return s.Body
}

func (s *CreateFunctionRequest) SetBody(v *CreateFunctionInput) *CreateFunctionRequest {
	s.Body = v
	return s
}

func (s *CreateFunctionRequest) Validate() error {
	return dara.Validate(s)
}
