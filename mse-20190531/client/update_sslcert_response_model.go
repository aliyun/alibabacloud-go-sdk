// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSSLCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSSLCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSSLCertResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSSLCertResponseBody) *UpdateSSLCertResponse
	GetBody() *UpdateSSLCertResponseBody
}

type UpdateSSLCertResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSSLCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSSLCertResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSSLCertResponse) GoString() string {
	return s.String()
}

func (s *UpdateSSLCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSSLCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSSLCertResponse) GetBody() *UpdateSSLCertResponseBody {
	return s.Body
}

func (s *UpdateSSLCertResponse) SetHeaders(v map[string]*string) *UpdateSSLCertResponse {
	s.Headers = v
	return s
}

func (s *UpdateSSLCertResponse) SetStatusCode(v int32) *UpdateSSLCertResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSSLCertResponse) SetBody(v *UpdateSSLCertResponseBody) *UpdateSSLCertResponse {
	s.Body = v
	return s
}

func (s *UpdateSSLCertResponse) Validate() error {
	return dara.Validate(s)
}
