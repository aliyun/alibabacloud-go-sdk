// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppRecordStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppRecordStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppRecordStatusResponseBody) *DescribeAppRecordStatusResponse
	GetBody() *DescribeAppRecordStatusResponseBody
}

type DescribeAppRecordStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppRecordStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppRecordStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppRecordStatusResponse) GetBody() *DescribeAppRecordStatusResponseBody {
	return s.Body
}

func (s *DescribeAppRecordStatusResponse) SetHeaders(v map[string]*string) *DescribeAppRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppRecordStatusResponse) SetStatusCode(v int32) *DescribeAppRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppRecordStatusResponse) SetBody(v *DescribeAppRecordStatusResponseBody) *DescribeAppRecordStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAppRecordStatusResponse) Validate() error {
	return dara.Validate(s)
}
