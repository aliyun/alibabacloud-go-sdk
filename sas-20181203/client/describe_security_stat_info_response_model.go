// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityStatInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityStatInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityStatInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityStatInfoResponseBody) *DescribeSecurityStatInfoResponse
	GetBody() *DescribeSecurityStatInfoResponseBody
}

type DescribeSecurityStatInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityStatInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityStatInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityStatInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityStatInfoResponse) GetBody() *DescribeSecurityStatInfoResponseBody {
	return s.Body
}

func (s *DescribeSecurityStatInfoResponse) SetHeaders(v map[string]*string) *DescribeSecurityStatInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityStatInfoResponse) SetStatusCode(v int32) *DescribeSecurityStatInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityStatInfoResponse) SetBody(v *DescribeSecurityStatInfoResponseBody) *DescribeSecurityStatInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityStatInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
