// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStreamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStreamsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStreamsResponseBody) *DescribeStreamsResponse
	GetBody() *DescribeStreamsResponseBody
}

type DescribeStreamsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStreamsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeStreamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStreamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStreamsResponse) GetBody() *DescribeStreamsResponseBody {
	return s.Body
}

func (s *DescribeStreamsResponse) SetHeaders(v map[string]*string) *DescribeStreamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeStreamsResponse) SetStatusCode(v int32) *DescribeStreamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStreamsResponse) SetBody(v *DescribeStreamsResponseBody) *DescribeStreamsResponse {
	s.Body = v
	return s
}

func (s *DescribeStreamsResponse) Validate() error {
	return dara.Validate(s)
}
