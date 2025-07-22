// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppRecordTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppRecordTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppRecordTemplatesResponseBody) *DescribeAppRecordTemplatesResponse
	GetBody() *DescribeAppRecordTemplatesResponseBody
}

type DescribeAppRecordTemplatesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppRecordTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppRecordTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppRecordTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppRecordTemplatesResponse) GetBody() *DescribeAppRecordTemplatesResponseBody {
	return s.Body
}

func (s *DescribeAppRecordTemplatesResponse) SetHeaders(v map[string]*string) *DescribeAppRecordTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppRecordTemplatesResponse) SetStatusCode(v int32) *DescribeAppRecordTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppRecordTemplatesResponse) SetBody(v *DescribeAppRecordTemplatesResponseBody) *DescribeAppRecordTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppRecordTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
