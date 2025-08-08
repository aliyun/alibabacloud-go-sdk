// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateShareResponse
	GetStatusCode() *int32
}

type CreateShareResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateShareResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateShareResponse) GoString() string {
	return s.String()
}

func (s *CreateShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateShareResponse) SetHeaders(v map[string]*string) *CreateShareResponse {
	s.Headers = v
	return s
}

func (s *CreateShareResponse) SetStatusCode(v int32) *CreateShareResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShareResponse) Validate() error {
	return dara.Validate(s)
}
