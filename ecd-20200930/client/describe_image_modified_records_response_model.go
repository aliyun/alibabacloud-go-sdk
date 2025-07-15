// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageModifiedRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageModifiedRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageModifiedRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageModifiedRecordsResponseBody) *DescribeImageModifiedRecordsResponse
	GetBody() *DescribeImageModifiedRecordsResponseBody
}

type DescribeImageModifiedRecordsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageModifiedRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageModifiedRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModifiedRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageModifiedRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageModifiedRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageModifiedRecordsResponse) GetBody() *DescribeImageModifiedRecordsResponseBody {
	return s.Body
}

func (s *DescribeImageModifiedRecordsResponse) SetHeaders(v map[string]*string) *DescribeImageModifiedRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageModifiedRecordsResponse) SetStatusCode(v int32) *DescribeImageModifiedRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageModifiedRecordsResponse) SetBody(v *DescribeImageModifiedRecordsResponseBody) *DescribeImageModifiedRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeImageModifiedRecordsResponse) Validate() error {
	return dara.Validate(s)
}
