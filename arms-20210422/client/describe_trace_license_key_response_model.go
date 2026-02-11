// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceLicenseKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTraceLicenseKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTraceLicenseKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTraceLicenseKeyResponseBody) *DescribeTraceLicenseKeyResponse
	GetBody() *DescribeTraceLicenseKeyResponseBody
}

type DescribeTraceLicenseKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTraceLicenseKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTraceLicenseKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTraceLicenseKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTraceLicenseKeyResponse) GetBody() *DescribeTraceLicenseKeyResponseBody {
	return s.Body
}

func (s *DescribeTraceLicenseKeyResponse) SetHeaders(v map[string]*string) *DescribeTraceLicenseKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) SetStatusCode(v int32) *DescribeTraceLicenseKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) SetBody(v *DescribeTraceLicenseKeyResponseBody) *DescribeTraceLicenseKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
