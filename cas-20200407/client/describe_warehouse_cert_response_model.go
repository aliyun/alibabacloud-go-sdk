// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarehouseCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWarehouseCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWarehouseCertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWarehouseCertResponseBody) *DescribeWarehouseCertResponse
	GetBody() *DescribeWarehouseCertResponseBody
}

type DescribeWarehouseCertResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWarehouseCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWarehouseCertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarehouseCertResponse) GoString() string {
	return s.String()
}

func (s *DescribeWarehouseCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWarehouseCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWarehouseCertResponse) GetBody() *DescribeWarehouseCertResponseBody {
	return s.Body
}

func (s *DescribeWarehouseCertResponse) SetHeaders(v map[string]*string) *DescribeWarehouseCertResponse {
	s.Headers = v
	return s
}

func (s *DescribeWarehouseCertResponse) SetStatusCode(v int32) *DescribeWarehouseCertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWarehouseCertResponse) SetBody(v *DescribeWarehouseCertResponseBody) *DescribeWarehouseCertResponse {
	s.Body = v
	return s
}

func (s *DescribeWarehouseCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
