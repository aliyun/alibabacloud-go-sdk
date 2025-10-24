// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBucketsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBucketsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBucketsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBucketsResponseBody) *DescribeBucketsResponse
	GetBody() *DescribeBucketsResponseBody
}

type DescribeBucketsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBucketsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBucketsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBucketsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBucketsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBucketsResponse) GetBody() *DescribeBucketsResponseBody {
	return s.Body
}

func (s *DescribeBucketsResponse) SetHeaders(v map[string]*string) *DescribeBucketsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBucketsResponse) SetStatusCode(v int32) *DescribeBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBucketsResponse) SetBody(v *DescribeBucketsResponseBody) *DescribeBucketsResponse {
	s.Body = v
	return s
}

func (s *DescribeBucketsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
