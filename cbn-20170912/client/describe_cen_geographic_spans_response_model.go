// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenGeographicSpansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenGeographicSpansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenGeographicSpansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenGeographicSpansResponseBody) *DescribeCenGeographicSpansResponse
	GetBody() *DescribeCenGeographicSpansResponseBody
}

type DescribeCenGeographicSpansResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenGeographicSpansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenGeographicSpansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenGeographicSpansResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenGeographicSpansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenGeographicSpansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenGeographicSpansResponse) GetBody() *DescribeCenGeographicSpansResponseBody {
	return s.Body
}

func (s *DescribeCenGeographicSpansResponse) SetHeaders(v map[string]*string) *DescribeCenGeographicSpansResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenGeographicSpansResponse) SetStatusCode(v int32) *DescribeCenGeographicSpansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenGeographicSpansResponse) SetBody(v *DescribeCenGeographicSpansResponseBody) *DescribeCenGeographicSpansResponse {
	s.Body = v
	return s
}

func (s *DescribeCenGeographicSpansResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
