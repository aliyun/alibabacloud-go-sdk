// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelServicesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelServicesResponseBody) *DescribeModelServicesResponse
	GetBody() *DescribeModelServicesResponseBody
}

type DescribeModelServicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelServicesResponse) GetBody() *DescribeModelServicesResponseBody {
	return s.Body
}

func (s *DescribeModelServicesResponse) SetHeaders(v map[string]*string) *DescribeModelServicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelServicesResponse) SetStatusCode(v int32) *DescribeModelServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelServicesResponse) SetBody(v *DescribeModelServicesResponseBody) *DescribeModelServicesResponse {
	s.Body = v
	return s
}

func (s *DescribeModelServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
