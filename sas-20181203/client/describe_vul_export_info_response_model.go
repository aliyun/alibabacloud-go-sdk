// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulExportInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulExportInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulExportInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulExportInfoResponseBody) *DescribeVulExportInfoResponse
	GetBody() *DescribeVulExportInfoResponseBody
}

type DescribeVulExportInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulExportInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulExportInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulExportInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulExportInfoResponse) GetBody() *DescribeVulExportInfoResponseBody {
	return s.Body
}

func (s *DescribeVulExportInfoResponse) SetHeaders(v map[string]*string) *DescribeVulExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulExportInfoResponse) SetStatusCode(v int32) *DescribeVulExportInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulExportInfoResponse) SetBody(v *DescribeVulExportInfoResponseBody) *DescribeVulExportInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeVulExportInfoResponse) Validate() error {
	return dara.Validate(s)
}
