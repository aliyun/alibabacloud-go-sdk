// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllEventNameAndCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllEventNameAndCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllEventNameAndCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllEventNameAndCodeResponseBody) *DescribeAllEventNameAndCodeResponse
	GetBody() *DescribeAllEventNameAndCodeResponseBody
}

type DescribeAllEventNameAndCodeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllEventNameAndCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllEventNameAndCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllEventNameAndCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllEventNameAndCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllEventNameAndCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllEventNameAndCodeResponse) GetBody() *DescribeAllEventNameAndCodeResponseBody {
	return s.Body
}

func (s *DescribeAllEventNameAndCodeResponse) SetHeaders(v map[string]*string) *DescribeAllEventNameAndCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllEventNameAndCodeResponse) SetStatusCode(v int32) *DescribeAllEventNameAndCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllEventNameAndCodeResponse) SetBody(v *DescribeAllEventNameAndCodeResponseBody) *DescribeAllEventNameAndCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeAllEventNameAndCodeResponse) Validate() error {
	return dara.Validate(s)
}
