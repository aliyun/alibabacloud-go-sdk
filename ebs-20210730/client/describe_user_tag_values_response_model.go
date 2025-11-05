// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagValuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserTagValuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserTagValuesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserTagValuesResponseBody) *DescribeUserTagValuesResponse
	GetBody() *DescribeUserTagValuesResponseBody
}

type DescribeUserTagValuesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserTagValuesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagValuesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTagValuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserTagValuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserTagValuesResponse) GetBody() *DescribeUserTagValuesResponseBody {
	return s.Body
}

func (s *DescribeUserTagValuesResponse) SetHeaders(v map[string]*string) *DescribeUserTagValuesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTagValuesResponse) SetStatusCode(v int32) *DescribeUserTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserTagValuesResponse) SetBody(v *DescribeUserTagValuesResponseBody) *DescribeUserTagValuesResponse {
	s.Body = v
	return s
}

func (s *DescribeUserTagValuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
