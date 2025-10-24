// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceResponse
	GetStatusCode() *int32
	SetBody(v *Service) *DescribeServiceResponse
	GetBody() *Service
}

type DescribeServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Service           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceResponse) GetBody() *Service {
	return s.Body
}

func (s *DescribeServiceResponse) SetHeaders(v map[string]*string) *DescribeServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceResponse) SetStatusCode(v int32) *DescribeServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceResponse) SetBody(v *Service) *DescribeServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
