// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVendorListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVendorListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVendorListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVendorListResponseBody) *DescribeVendorListResponse
	GetBody() *DescribeVendorListResponseBody
}

type DescribeVendorListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVendorListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVendorListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVendorListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVendorListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVendorListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVendorListResponse) GetBody() *DescribeVendorListResponseBody {
	return s.Body
}

func (s *DescribeVendorListResponse) SetHeaders(v map[string]*string) *DescribeVendorListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVendorListResponse) SetStatusCode(v int32) *DescribeVendorListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVendorListResponse) SetBody(v *DescribeVendorListResponseBody) *DescribeVendorListResponse {
	s.Body = v
	return s
}

func (s *DescribeVendorListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
