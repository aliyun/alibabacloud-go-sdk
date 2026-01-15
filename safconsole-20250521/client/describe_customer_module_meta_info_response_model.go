// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleMetaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomerModuleMetaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomerModuleMetaInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomerModuleMetaInfoResponseBody) *DescribeCustomerModuleMetaInfoResponse
	GetBody() *DescribeCustomerModuleMetaInfoResponseBody
}

type DescribeCustomerModuleMetaInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomerModuleMetaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomerModuleMetaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleMetaInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleMetaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomerModuleMetaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomerModuleMetaInfoResponse) GetBody() *DescribeCustomerModuleMetaInfoResponseBody {
	return s.Body
}

func (s *DescribeCustomerModuleMetaInfoResponse) SetHeaders(v map[string]*string) *DescribeCustomerModuleMetaInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponse) SetStatusCode(v int32) *DescribeCustomerModuleMetaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponse) SetBody(v *DescribeCustomerModuleMetaInfoResponseBody) *DescribeCustomerModuleMetaInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomerModuleMetaInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
