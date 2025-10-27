// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstallCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstallCodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstallCodeResponseBody) *DescribeInstallCodeResponse
	GetBody() *DescribeInstallCodeResponseBody
}

type DescribeInstallCodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstallCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstallCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstallCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstallCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstallCodeResponse) GetBody() *DescribeInstallCodeResponseBody {
	return s.Body
}

func (s *DescribeInstallCodeResponse) SetHeaders(v map[string]*string) *DescribeInstallCodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstallCodeResponse) SetStatusCode(v int32) *DescribeInstallCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstallCodeResponse) SetBody(v *DescribeInstallCodeResponseBody) *DescribeInstallCodeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstallCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
