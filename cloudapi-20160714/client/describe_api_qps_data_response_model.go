// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiQpsDataResponseBody) *DescribeApiQpsDataResponse
	GetBody() *DescribeApiQpsDataResponseBody
}

type DescribeApiQpsDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiQpsDataResponse) GetBody() *DescribeApiQpsDataResponseBody {
	return s.Body
}

func (s *DescribeApiQpsDataResponse) SetHeaders(v map[string]*string) *DescribeApiQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiQpsDataResponse) SetStatusCode(v int32) *DescribeApiQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiQpsDataResponse) SetBody(v *DescribeApiQpsDataResponseBody) *DescribeApiQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeApiQpsDataResponse) Validate() error {
	return dara.Validate(s)
}
