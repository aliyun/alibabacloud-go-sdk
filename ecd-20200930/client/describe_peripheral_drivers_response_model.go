// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePeripheralDriversResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePeripheralDriversResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePeripheralDriversResponse
	GetStatusCode() *int32
	SetBody(v *DescribePeripheralDriversResponseBody) *DescribePeripheralDriversResponse
	GetBody() *DescribePeripheralDriversResponseBody
}

type DescribePeripheralDriversResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePeripheralDriversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePeripheralDriversResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePeripheralDriversResponse) GoString() string {
	return s.String()
}

func (s *DescribePeripheralDriversResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePeripheralDriversResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePeripheralDriversResponse) GetBody() *DescribePeripheralDriversResponseBody {
	return s.Body
}

func (s *DescribePeripheralDriversResponse) SetHeaders(v map[string]*string) *DescribePeripheralDriversResponse {
	s.Headers = v
	return s
}

func (s *DescribePeripheralDriversResponse) SetStatusCode(v int32) *DescribePeripheralDriversResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePeripheralDriversResponse) SetBody(v *DescribePeripheralDriversResponseBody) *DescribePeripheralDriversResponse {
	s.Body = v
	return s
}

func (s *DescribePeripheralDriversResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
