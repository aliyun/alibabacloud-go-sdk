// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCertificateStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCertificateStateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCertificateStateResponseBody) *DescribeCertificateStateResponse
	GetBody() *DescribeCertificateStateResponseBody
}

type DescribeCertificateStateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificateStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificateStateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateStateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCertificateStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCertificateStateResponse) GetBody() *DescribeCertificateStateResponseBody {
	return s.Body
}

func (s *DescribeCertificateStateResponse) SetHeaders(v map[string]*string) *DescribeCertificateStateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateStateResponse) SetStatusCode(v int32) *DescribeCertificateStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificateStateResponse) SetBody(v *DescribeCertificateStateResponseBody) *DescribeCertificateStateResponse {
	s.Body = v
	return s
}

func (s *DescribeCertificateStateResponse) Validate() error {
	return dara.Validate(s)
}
