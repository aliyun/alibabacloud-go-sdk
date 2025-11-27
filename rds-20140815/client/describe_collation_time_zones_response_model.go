// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollationTimeZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCollationTimeZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCollationTimeZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCollationTimeZonesResponseBody) *DescribeCollationTimeZonesResponse
	GetBody() *DescribeCollationTimeZonesResponseBody
}

type DescribeCollationTimeZonesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCollationTimeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCollationTimeZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollationTimeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCollationTimeZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCollationTimeZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCollationTimeZonesResponse) GetBody() *DescribeCollationTimeZonesResponseBody {
	return s.Body
}

func (s *DescribeCollationTimeZonesResponse) SetHeaders(v map[string]*string) *DescribeCollationTimeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCollationTimeZonesResponse) SetStatusCode(v int32) *DescribeCollationTimeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCollationTimeZonesResponse) SetBody(v *DescribeCollationTimeZonesResponseBody) *DescribeCollationTimeZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeCollationTimeZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
