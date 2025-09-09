// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecycleBinStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecycleBinStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecycleBinStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecycleBinStatusResponseBody) *DescribeRecycleBinStatusResponse
	GetBody() *DescribeRecycleBinStatusResponseBody
}

type DescribeRecycleBinStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecycleBinStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecycleBinStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecycleBinStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecycleBinStatusResponse) GetBody() *DescribeRecycleBinStatusResponseBody {
	return s.Body
}

func (s *DescribeRecycleBinStatusResponse) SetHeaders(v map[string]*string) *DescribeRecycleBinStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecycleBinStatusResponse) SetStatusCode(v int32) *DescribeRecycleBinStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecycleBinStatusResponse) SetBody(v *DescribeRecycleBinStatusResponseBody) *DescribeRecycleBinStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRecycleBinStatusResponse) Validate() error {
	return dara.Validate(s)
}
