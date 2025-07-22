// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRecordTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRecordTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRecordTemplatesResponseBody) *DescribeRecordTemplatesResponse
	GetBody() *DescribeRecordTemplatesResponseBody
}

type DescribeRecordTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRecordTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRecordTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRecordTemplatesResponse) GetBody() *DescribeRecordTemplatesResponseBody {
	return s.Body
}

func (s *DescribeRecordTemplatesResponse) SetHeaders(v map[string]*string) *DescribeRecordTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordTemplatesResponse) SetStatusCode(v int32) *DescribeRecordTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordTemplatesResponse) SetBody(v *DescribeRecordTemplatesResponseBody) *DescribeRecordTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeRecordTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
