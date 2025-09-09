// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecursionRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecursionRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecursionRecordResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecursionRecordResponseBody) *DescribeRecursionRecordResponse
	GetBody() *DescribeRecursionRecordResponseBody
}

type DescribeRecursionRecordResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecursionRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecursionRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecursionRecordResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecursionRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecursionRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecursionRecordResponse) GetBody() *DescribeRecursionRecordResponseBody {
	return s.Body
}

func (s *DescribeRecursionRecordResponse) SetHeaders(v map[string]*string) *DescribeRecursionRecordResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecursionRecordResponse) SetStatusCode(v int32) *DescribeRecursionRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecursionRecordResponse) SetBody(v *DescribeRecursionRecordResponseBody) *DescribeRecursionRecordResponse {
	s.Body = v
	return s
}

func (s *DescribeRecursionRecordResponse) Validate() error {
	return dara.Validate(s)
}
