// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleOutputInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomerModuleOutputInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomerModuleOutputInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomerModuleOutputInfoResponseBody) *DescribeCustomerModuleOutputInfoResponse
	GetBody() *DescribeCustomerModuleOutputInfoResponseBody
}

type DescribeCustomerModuleOutputInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomerModuleOutputInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomerModuleOutputInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleOutputInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleOutputInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomerModuleOutputInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomerModuleOutputInfoResponse) GetBody() *DescribeCustomerModuleOutputInfoResponseBody {
	return s.Body
}

func (s *DescribeCustomerModuleOutputInfoResponse) SetHeaders(v map[string]*string) *DescribeCustomerModuleOutputInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponse) SetStatusCode(v int32) *DescribeCustomerModuleOutputInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponse) SetBody(v *DescribeCustomerModuleOutputInfoResponseBody) *DescribeCustomerModuleOutputInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
