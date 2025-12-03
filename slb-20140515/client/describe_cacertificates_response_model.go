// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCACertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCACertificatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCACertificatesResponseBody) *DescribeCACertificatesResponse
	GetBody() *DescribeCACertificatesResponseBody
}

type DescribeCACertificatesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCACertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCACertificatesResponse) GetBody() *DescribeCACertificatesResponseBody {
	return s.Body
}

func (s *DescribeCACertificatesResponse) SetHeaders(v map[string]*string) *DescribeCACertificatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificatesResponse) SetStatusCode(v int32) *DescribeCACertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificatesResponse) SetBody(v *DescribeCACertificatesResponseBody) *DescribeCACertificatesResponse {
	s.Body = v
	return s
}

func (s *DescribeCACertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
