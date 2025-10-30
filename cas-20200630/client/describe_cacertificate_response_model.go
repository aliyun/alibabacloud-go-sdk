// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCACertificateResponseBody) *DescribeCACertificateResponse
	GetBody() *DescribeCACertificateResponseBody
}

type DescribeCACertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCACertificateResponse) GetBody() *DescribeCACertificateResponseBody {
	return s.Body
}

func (s *DescribeCACertificateResponse) SetHeaders(v map[string]*string) *DescribeCACertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificateResponse) SetStatusCode(v int32) *DescribeCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificateResponse) SetBody(v *DescribeCACertificateResponseBody) *DescribeCACertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeCACertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
