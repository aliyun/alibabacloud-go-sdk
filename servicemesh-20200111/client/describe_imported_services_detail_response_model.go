// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportedServicesDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImportedServicesDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImportedServicesDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImportedServicesDetailResponseBody) *DescribeImportedServicesDetailResponse
	GetBody() *DescribeImportedServicesDetailResponseBody
}

type DescribeImportedServicesDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImportedServicesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImportedServicesDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportedServicesDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportedServicesDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImportedServicesDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImportedServicesDetailResponse) GetBody() *DescribeImportedServicesDetailResponseBody {
	return s.Body
}

func (s *DescribeImportedServicesDetailResponse) SetHeaders(v map[string]*string) *DescribeImportedServicesDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportedServicesDetailResponse) SetStatusCode(v int32) *DescribeImportedServicesDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportedServicesDetailResponse) SetBody(v *DescribeImportedServicesDetailResponseBody) *DescribeImportedServicesDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeImportedServicesDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
