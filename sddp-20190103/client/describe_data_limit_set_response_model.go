// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataLimitSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataLimitSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataLimitSetResponseBody) *DescribeDataLimitSetResponse
	GetBody() *DescribeDataLimitSetResponseBody
}

type DescribeDataLimitSetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataLimitSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataLimitSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataLimitSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataLimitSetResponse) GetBody() *DescribeDataLimitSetResponseBody {
	return s.Body
}

func (s *DescribeDataLimitSetResponse) SetHeaders(v map[string]*string) *DescribeDataLimitSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataLimitSetResponse) SetStatusCode(v int32) *DescribeDataLimitSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataLimitSetResponse) SetBody(v *DescribeDataLimitSetResponseBody) *DescribeDataLimitSetResponse {
	s.Body = v
	return s
}

func (s *DescribeDataLimitSetResponse) Validate() error {
	return dara.Validate(s)
}
