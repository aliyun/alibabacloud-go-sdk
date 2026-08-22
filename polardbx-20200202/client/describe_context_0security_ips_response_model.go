// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0SecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContext0SecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContext0SecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContext0SecurityIpsResponseBody) *DescribeContext0SecurityIpsResponse
	GetBody() *DescribeContext0SecurityIpsResponseBody
}

type DescribeContext0SecurityIpsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContext0SecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContext0SecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0SecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContext0SecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContext0SecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContext0SecurityIpsResponse) GetBody() *DescribeContext0SecurityIpsResponseBody {
	return s.Body
}

func (s *DescribeContext0SecurityIpsResponse) SetHeaders(v map[string]*string) *DescribeContext0SecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContext0SecurityIpsResponse) SetStatusCode(v int32) *DescribeContext0SecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContext0SecurityIpsResponse) SetBody(v *DescribeContext0SecurityIpsResponseBody) *DescribeContext0SecurityIpsResponse {
	s.Body = v
	return s
}

func (s *DescribeContext0SecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
