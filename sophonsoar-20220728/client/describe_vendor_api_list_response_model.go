// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVendorApiListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVendorApiListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVendorApiListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVendorApiListResponseBody) *DescribeVendorApiListResponse
	GetBody() *DescribeVendorApiListResponseBody
}

type DescribeVendorApiListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVendorApiListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVendorApiListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVendorApiListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVendorApiListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVendorApiListResponse) GetBody() *DescribeVendorApiListResponseBody {
	return s.Body
}

func (s *DescribeVendorApiListResponse) SetHeaders(v map[string]*string) *DescribeVendorApiListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVendorApiListResponse) SetStatusCode(v int32) *DescribeVendorApiListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVendorApiListResponse) SetBody(v *DescribeVendorApiListResponseBody) *DescribeVendorApiListResponse {
	s.Body = v
	return s
}

func (s *DescribeVendorApiListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
