// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityModeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityModeResponseBody) *DescribeSecurityModeResponse
	GetBody() *DescribeSecurityModeResponseBody
}

type DescribeSecurityModeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityModeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityModeResponse) GetBody() *DescribeSecurityModeResponseBody {
	return s.Body
}

func (s *DescribeSecurityModeResponse) SetHeaders(v map[string]*string) *DescribeSecurityModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityModeResponse) SetStatusCode(v int32) *DescribeSecurityModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityModeResponse) SetBody(v *DescribeSecurityModeResponseBody) *DescribeSecurityModeResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
