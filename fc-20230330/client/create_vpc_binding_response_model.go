// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcBindingResponse
	GetStatusCode() *int32
}

type CreateVpcBindingResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateVpcBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcBindingResponse) SetHeaders(v map[string]*string) *CreateVpcBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcBindingResponse) SetStatusCode(v int32) *CreateVpcBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcBindingResponse) Validate() error {
	return dara.Validate(s)
}
