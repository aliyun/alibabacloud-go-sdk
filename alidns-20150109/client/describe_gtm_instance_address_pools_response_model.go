// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceAddressPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmInstanceAddressPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmInstanceAddressPoolsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmInstanceAddressPoolsResponseBody) *DescribeGtmInstanceAddressPoolsResponse
	GetBody() *DescribeGtmInstanceAddressPoolsResponseBody
}

type DescribeGtmInstanceAddressPoolsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmInstanceAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmInstanceAddressPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmInstanceAddressPoolsResponse) GetBody() *DescribeGtmInstanceAddressPoolsResponseBody {
	return s.Body
}

func (s *DescribeGtmInstanceAddressPoolsResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponse) SetStatusCode(v int32) *DescribeGtmInstanceAddressPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponse) SetBody(v *DescribeGtmInstanceAddressPoolsResponseBody) *DescribeGtmInstanceAddressPoolsResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
