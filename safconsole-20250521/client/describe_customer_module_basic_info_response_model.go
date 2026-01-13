// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomerModuleBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomerModuleBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomerModuleBasicInfoResponseBody) *DescribeCustomerModuleBasicInfoResponse
	GetBody() *DescribeCustomerModuleBasicInfoResponseBody
}

type DescribeCustomerModuleBasicInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomerModuleBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomerModuleBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomerModuleBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomerModuleBasicInfoResponse) GetBody() *DescribeCustomerModuleBasicInfoResponseBody {
	return s.Body
}

func (s *DescribeCustomerModuleBasicInfoResponse) SetHeaders(v map[string]*string) *DescribeCustomerModuleBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponse) SetStatusCode(v int32) *DescribeCustomerModuleBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponse) SetBody(v *DescribeCustomerModuleBasicInfoResponseBody) *DescribeCustomerModuleBasicInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
