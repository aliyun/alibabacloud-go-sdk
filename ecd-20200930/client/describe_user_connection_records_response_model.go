// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConnectionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserConnectionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserConnectionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserConnectionRecordsResponseBody) *DescribeUserConnectionRecordsResponse
	GetBody() *DescribeUserConnectionRecordsResponseBody
}

type DescribeUserConnectionRecordsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserConnectionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserConnectionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserConnectionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserConnectionRecordsResponse) GetBody() *DescribeUserConnectionRecordsResponseBody {
	return s.Body
}

func (s *DescribeUserConnectionRecordsResponse) SetHeaders(v map[string]*string) *DescribeUserConnectionRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) SetStatusCode(v int32) *DescribeUserConnectionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) SetBody(v *DescribeUserConnectionRecordsResponseBody) *DescribeUserConnectionRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserConnectionRecordsResponse) Validate() error {
	return dara.Validate(s)
}
