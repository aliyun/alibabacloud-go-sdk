// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCertificateByCaSnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeviceCertificateByCaSnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeviceCertificateByCaSnResponse
	GetStatusCode() *int32
	SetBody(v *ListDeviceCertificateByCaSnResponseBody) *ListDeviceCertificateByCaSnResponse
	GetBody() *ListDeviceCertificateByCaSnResponseBody
}

type ListDeviceCertificateByCaSnResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeviceCertificateByCaSnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeviceCertificateByCaSnResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCertificateByCaSnResponse) GoString() string {
	return s.String()
}

func (s *ListDeviceCertificateByCaSnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeviceCertificateByCaSnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeviceCertificateByCaSnResponse) GetBody() *ListDeviceCertificateByCaSnResponseBody {
	return s.Body
}

func (s *ListDeviceCertificateByCaSnResponse) SetHeaders(v map[string]*string) *ListDeviceCertificateByCaSnResponse {
	s.Headers = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponse) SetStatusCode(v int32) *ListDeviceCertificateByCaSnResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeviceCertificateByCaSnResponse) SetBody(v *ListDeviceCertificateByCaSnResponseBody) *ListDeviceCertificateByCaSnResponse {
	s.Body = v
	return s
}

func (s *ListDeviceCertificateByCaSnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
