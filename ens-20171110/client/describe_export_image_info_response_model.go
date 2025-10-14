// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportImageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExportImageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExportImageInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExportImageInfoResponseBody) *DescribeExportImageInfoResponse
	GetBody() *DescribeExportImageInfoResponseBody
}

type DescribeExportImageInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExportImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExportImageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportImageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportImageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExportImageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExportImageInfoResponse) GetBody() *DescribeExportImageInfoResponseBody {
	return s.Body
}

func (s *DescribeExportImageInfoResponse) SetHeaders(v map[string]*string) *DescribeExportImageInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportImageInfoResponse) SetStatusCode(v int32) *DescribeExportImageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExportImageInfoResponse) SetBody(v *DescribeExportImageInfoResponseBody) *DescribeExportImageInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeExportImageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
