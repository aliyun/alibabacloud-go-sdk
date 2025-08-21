// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportImageStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExportImageStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExportImageStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExportImageStatusResponseBody) *DescribeExportImageStatusResponse
	GetBody() *DescribeExportImageStatusResponseBody
}

type DescribeExportImageStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExportImageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExportImageStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportImageStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExportImageStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExportImageStatusResponse) GetBody() *DescribeExportImageStatusResponseBody {
	return s.Body
}

func (s *DescribeExportImageStatusResponse) SetHeaders(v map[string]*string) *DescribeExportImageStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportImageStatusResponse) SetStatusCode(v int32) *DescribeExportImageStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExportImageStatusResponse) SetBody(v *DescribeExportImageStatusResponseBody) *DescribeExportImageStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeExportImageStatusResponse) Validate() error {
	return dara.Validate(s)
}
