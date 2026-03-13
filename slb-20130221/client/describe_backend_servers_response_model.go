// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackendServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackendServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackendServersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackendServersResponseBody) *DescribeBackendServersResponse
	GetBody() *DescribeBackendServersResponseBody
}

type DescribeBackendServersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackendServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackendServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackendServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackendServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackendServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackendServersResponse) GetBody() *DescribeBackendServersResponseBody {
	return s.Body
}

func (s *DescribeBackendServersResponse) SetHeaders(v map[string]*string) *DescribeBackendServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackendServersResponse) SetStatusCode(v int32) *DescribeBackendServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackendServersResponse) SetBody(v *DescribeBackendServersResponseBody) *DescribeBackendServersResponse {
	s.Body = v
	return s
}

func (s *DescribeBackendServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
