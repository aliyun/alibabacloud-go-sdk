// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExportInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExportInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExportInfoResponseBody) *DescribeExportInfoResponse
	GetBody() *DescribeExportInfoResponseBody
}

type DescribeExportInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExportInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExportInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExportInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExportInfoResponse) GetBody() *DescribeExportInfoResponseBody {
	return s.Body
}

func (s *DescribeExportInfoResponse) SetHeaders(v map[string]*string) *DescribeExportInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportInfoResponse) SetStatusCode(v int32) *DescribeExportInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExportInfoResponse) SetBody(v *DescribeExportInfoResponseBody) *DescribeExportInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeExportInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
