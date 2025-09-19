// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogMetaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogMetaResponseBody) *DescribeLogMetaResponse
	GetBody() *DescribeLogMetaResponseBody
}

type DescribeLogMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogMetaResponse) GetBody() *DescribeLogMetaResponseBody {
	return s.Body
}

func (s *DescribeLogMetaResponse) SetHeaders(v map[string]*string) *DescribeLogMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogMetaResponse) SetStatusCode(v int32) *DescribeLogMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogMetaResponse) SetBody(v *DescribeLogMetaResponseBody) *DescribeLogMetaResponse {
	s.Body = v
	return s
}

func (s *DescribeLogMetaResponse) Validate() error {
	return dara.Validate(s)
}
