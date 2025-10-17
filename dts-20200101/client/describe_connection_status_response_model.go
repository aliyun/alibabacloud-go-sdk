// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConnectionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConnectionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConnectionStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConnectionStatusResponseBody) *DescribeConnectionStatusResponse
	GetBody() *DescribeConnectionStatusResponseBody
}

type DescribeConnectionStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConnectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConnectionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConnectionStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConnectionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConnectionStatusResponse) GetBody() *DescribeConnectionStatusResponseBody {
	return s.Body
}

func (s *DescribeConnectionStatusResponse) SetHeaders(v map[string]*string) *DescribeConnectionStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectionStatusResponse) SetStatusCode(v int32) *DescribeConnectionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConnectionStatusResponse) SetBody(v *DescribeConnectionStatusResponseBody) *DescribeConnectionStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeConnectionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
