// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagValuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagValuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagValuesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagValuesResponseBody) *DescribeTagValuesResponse
	GetBody() *DescribeTagValuesResponseBody
}

type DescribeTagValuesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagValuesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValuesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagValuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagValuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagValuesResponse) GetBody() *DescribeTagValuesResponseBody {
	return s.Body
}

func (s *DescribeTagValuesResponse) SetHeaders(v map[string]*string) *DescribeTagValuesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagValuesResponse) SetStatusCode(v int32) *DescribeTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagValuesResponse) SetBody(v *DescribeTagValuesResponseBody) *DescribeTagValuesResponse {
	s.Body = v
	return s
}

func (s *DescribeTagValuesResponse) Validate() error {
	return dara.Validate(s)
}
