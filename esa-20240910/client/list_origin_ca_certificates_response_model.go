// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginCaCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOriginCaCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOriginCaCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListOriginCaCertificatesResponseBody) *ListOriginCaCertificatesResponse
	GetBody() *ListOriginCaCertificatesResponseBody
}

type ListOriginCaCertificatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOriginCaCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOriginCaCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOriginCaCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListOriginCaCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOriginCaCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOriginCaCertificatesResponse) GetBody() *ListOriginCaCertificatesResponseBody {
	return s.Body
}

func (s *ListOriginCaCertificatesResponse) SetHeaders(v map[string]*string) *ListOriginCaCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListOriginCaCertificatesResponse) SetStatusCode(v int32) *ListOriginCaCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOriginCaCertificatesResponse) SetBody(v *ListOriginCaCertificatesResponseBody) *ListOriginCaCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListOriginCaCertificatesResponse) Validate() error {
	return dara.Validate(s)
}
