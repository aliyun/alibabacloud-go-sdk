// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcConfigResponse
	GetStatusCode() *int32
}

type UpdateVpcConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateVpcConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcConfigResponse) SetHeaders(v map[string]*string) *UpdateVpcConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcConfigResponse) SetStatusCode(v int32) *UpdateVpcConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcConfigResponse) Validate() error {
	return dara.Validate(s)
}
