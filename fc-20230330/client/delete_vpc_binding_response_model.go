// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcBindingResponse
	GetStatusCode() *int32
}

type DeleteVpcBindingResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteVpcBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcBindingResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcBindingResponse) SetHeaders(v map[string]*string) *DeleteVpcBindingResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcBindingResponse) SetStatusCode(v int32) *DeleteVpcBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcBindingResponse) Validate() error {
	return dara.Validate(s)
}
