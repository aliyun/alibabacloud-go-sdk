// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResultExportConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResultExportConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResultExportConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResultExportConfigResponseBody) *DescribeResultExportConfigResponse
	GetBody() *DescribeResultExportConfigResponseBody
}

type DescribeResultExportConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResultExportConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResultExportConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultExportConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeResultExportConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResultExportConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResultExportConfigResponse) GetBody() *DescribeResultExportConfigResponseBody {
	return s.Body
}

func (s *DescribeResultExportConfigResponse) SetHeaders(v map[string]*string) *DescribeResultExportConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeResultExportConfigResponse) SetStatusCode(v int32) *DescribeResultExportConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResultExportConfigResponse) SetBody(v *DescribeResultExportConfigResponseBody) *DescribeResultExportConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeResultExportConfigResponse) Validate() error {
	return dara.Validate(s)
}
