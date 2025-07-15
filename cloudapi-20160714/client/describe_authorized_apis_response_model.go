// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthorizedApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthorizedApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthorizedApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthorizedApisResponseBody) *DescribeAuthorizedApisResponse
	GetBody() *DescribeAuthorizedApisResponseBody
}

type DescribeAuthorizedApisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthorizedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthorizedApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthorizedApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthorizedApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthorizedApisResponse) GetBody() *DescribeAuthorizedApisResponseBody {
	return s.Body
}

func (s *DescribeAuthorizedApisResponse) SetHeaders(v map[string]*string) *DescribeAuthorizedApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthorizedApisResponse) SetStatusCode(v int32) *DescribeAuthorizedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthorizedApisResponse) SetBody(v *DescribeAuthorizedApisResponseBody) *DescribeAuthorizedApisResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthorizedApisResponse) Validate() error {
	return dara.Validate(s)
}
