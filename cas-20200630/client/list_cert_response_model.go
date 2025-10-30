// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCertResponse
	GetStatusCode() *int32
	SetBody(v *ListCertResponseBody) *ListCertResponse
	GetBody() *ListCertResponseBody
}

type ListCertResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponse) GoString() string {
	return s.String()
}

func (s *ListCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCertResponse) GetBody() *ListCertResponseBody {
	return s.Body
}

func (s *ListCertResponse) SetHeaders(v map[string]*string) *ListCertResponse {
	s.Headers = v
	return s
}

func (s *ListCertResponse) SetStatusCode(v int32) *ListCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCertResponse) SetBody(v *ListCertResponseBody) *ListCertResponse {
	s.Body = v
	return s
}

func (s *ListCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
