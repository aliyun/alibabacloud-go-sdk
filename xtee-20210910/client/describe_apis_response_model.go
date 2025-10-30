// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisResponseBody) *DescribeApisResponse
	GetBody() *DescribeApisResponseBody
}

type DescribeApisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisResponse) GetBody() *DescribeApisResponseBody {
	return s.Body
}

func (s *DescribeApisResponse) SetHeaders(v map[string]*string) *DescribeApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisResponse) SetStatusCode(v int32) *DescribeApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisResponse) SetBody(v *DescribeApisResponseBody) *DescribeApisResponse {
	s.Body = v
	return s
}

func (s *DescribeApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
