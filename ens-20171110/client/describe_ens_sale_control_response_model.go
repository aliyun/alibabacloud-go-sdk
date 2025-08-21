// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsSaleControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsSaleControlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsSaleControlResponseBody) *DescribeEnsSaleControlResponse
	GetBody() *DescribeEnsSaleControlResponseBody
}

type DescribeEnsSaleControlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsSaleControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsSaleControlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsSaleControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsSaleControlResponse) GetBody() *DescribeEnsSaleControlResponseBody {
	return s.Body
}

func (s *DescribeEnsSaleControlResponse) SetHeaders(v map[string]*string) *DescribeEnsSaleControlResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsSaleControlResponse) SetStatusCode(v int32) *DescribeEnsSaleControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsSaleControlResponse) SetBody(v *DescribeEnsSaleControlResponseBody) *DescribeEnsSaleControlResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsSaleControlResponse) Validate() error {
	return dara.Validate(s)
}
