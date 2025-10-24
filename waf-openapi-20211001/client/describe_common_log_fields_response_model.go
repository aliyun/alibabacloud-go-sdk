// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonLogFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonLogFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonLogFieldsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonLogFieldsResponseBody) *DescribeCommonLogFieldsResponse
	GetBody() *DescribeCommonLogFieldsResponseBody
}

type DescribeCommonLogFieldsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonLogFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonLogFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonLogFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonLogFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonLogFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonLogFieldsResponse) GetBody() *DescribeCommonLogFieldsResponseBody {
	return s.Body
}

func (s *DescribeCommonLogFieldsResponse) SetHeaders(v map[string]*string) *DescribeCommonLogFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonLogFieldsResponse) SetStatusCode(v int32) *DescribeCommonLogFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonLogFieldsResponse) SetBody(v *DescribeCommonLogFieldsResponseBody) *DescribeCommonLogFieldsResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonLogFieldsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
