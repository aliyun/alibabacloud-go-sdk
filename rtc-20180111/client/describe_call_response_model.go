// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCallResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCallResponseBody) *DescribeCallResponse
	GetBody() *DescribeCallResponseBody
}

type DescribeCallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCallResponse) GetBody() *DescribeCallResponseBody {
	return s.Body
}

func (s *DescribeCallResponse) SetHeaders(v map[string]*string) *DescribeCallResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallResponse) SetStatusCode(v int32) *DescribeCallResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallResponse) SetBody(v *DescribeCallResponseBody) *DescribeCallResponse {
	s.Body = v
	return s
}

func (s *DescribeCallResponse) Validate() error {
	return dara.Validate(s)
}
