// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmInstanceAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmInstanceAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmInstanceAddressPoolResponseBody) *DescribeGtmInstanceAddressPoolResponse
	GetBody() *DescribeGtmInstanceAddressPoolResponseBody
}

type DescribeGtmInstanceAddressPoolResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmInstanceAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmInstanceAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmInstanceAddressPoolResponse) GetBody() *DescribeGtmInstanceAddressPoolResponseBody {
	return s.Body
}

func (s *DescribeGtmInstanceAddressPoolResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponse) SetStatusCode(v int32) *DescribeGtmInstanceAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponse) SetBody(v *DescribeGtmInstanceAddressPoolResponseBody) *DescribeGtmInstanceAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
