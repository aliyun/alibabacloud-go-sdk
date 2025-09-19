// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrictEventNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStrictEventNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStrictEventNameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStrictEventNameResponseBody) *DescribeStrictEventNameResponse
	GetBody() *DescribeStrictEventNameResponseBody
}

type DescribeStrictEventNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStrictEventNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStrictEventNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrictEventNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeStrictEventNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStrictEventNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStrictEventNameResponse) GetBody() *DescribeStrictEventNameResponseBody {
	return s.Body
}

func (s *DescribeStrictEventNameResponse) SetHeaders(v map[string]*string) *DescribeStrictEventNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeStrictEventNameResponse) SetStatusCode(v int32) *DescribeStrictEventNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStrictEventNameResponse) SetBody(v *DescribeStrictEventNameResponseBody) *DescribeStrictEventNameResponse {
	s.Body = v
	return s
}

func (s *DescribeStrictEventNameResponse) Validate() error {
	return dara.Validate(s)
}
