// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsedServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsedServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsedServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsedServiceResponseBody) *DescribeUsedServiceResponse
	GetBody() *DescribeUsedServiceResponseBody
}

type DescribeUsedServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsedServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsedServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsedServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsedServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsedServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsedServiceResponse) GetBody() *DescribeUsedServiceResponseBody {
	return s.Body
}

func (s *DescribeUsedServiceResponse) SetHeaders(v map[string]*string) *DescribeUsedServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsedServiceResponse) SetStatusCode(v int32) *DescribeUsedServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsedServiceResponse) SetBody(v *DescribeUsedServiceResponseBody) *DescribeUsedServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeUsedServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
