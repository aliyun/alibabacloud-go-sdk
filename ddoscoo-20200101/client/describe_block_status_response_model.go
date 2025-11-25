// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlockStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBlockStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBlockStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBlockStatusResponseBody) *DescribeBlockStatusResponse
	GetBody() *DescribeBlockStatusResponseBody
}

type DescribeBlockStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBlockStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBlockStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlockStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBlockStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBlockStatusResponse) GetBody() *DescribeBlockStatusResponseBody {
	return s.Body
}

func (s *DescribeBlockStatusResponse) SetHeaders(v map[string]*string) *DescribeBlockStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlockStatusResponse) SetStatusCode(v int32) *DescribeBlockStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlockStatusResponse) SetBody(v *DescribeBlockStatusResponseBody) *DescribeBlockStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeBlockStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
