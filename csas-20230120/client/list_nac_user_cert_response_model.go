// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacUserCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNacUserCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNacUserCertResponse
	GetStatusCode() *int32
	SetBody(v *ListNacUserCertResponseBody) *ListNacUserCertResponse
	GetBody() *ListNacUserCertResponseBody
}

type ListNacUserCertResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNacUserCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNacUserCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNacUserCertResponse) GoString() string {
	return s.String()
}

func (s *ListNacUserCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNacUserCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNacUserCertResponse) GetBody() *ListNacUserCertResponseBody {
	return s.Body
}

func (s *ListNacUserCertResponse) SetHeaders(v map[string]*string) *ListNacUserCertResponse {
	s.Headers = v
	return s
}

func (s *ListNacUserCertResponse) SetStatusCode(v int32) *ListNacUserCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNacUserCertResponse) SetBody(v *ListNacUserCertResponseBody) *ListNacUserCertResponse {
	s.Body = v
	return s
}

func (s *ListNacUserCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
