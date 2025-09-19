// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarningExportInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWarningExportInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWarningExportInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWarningExportInfoResponseBody) *DescribeWarningExportInfoResponse
	GetBody() *DescribeWarningExportInfoResponseBody
}

type DescribeWarningExportInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWarningExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWarningExportInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarningExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeWarningExportInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWarningExportInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWarningExportInfoResponse) GetBody() *DescribeWarningExportInfoResponseBody {
	return s.Body
}

func (s *DescribeWarningExportInfoResponse) SetHeaders(v map[string]*string) *DescribeWarningExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeWarningExportInfoResponse) SetStatusCode(v int32) *DescribeWarningExportInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWarningExportInfoResponse) SetBody(v *DescribeWarningExportInfoResponseBody) *DescribeWarningExportInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeWarningExportInfoResponse) Validate() error {
	return dara.Validate(s)
}
