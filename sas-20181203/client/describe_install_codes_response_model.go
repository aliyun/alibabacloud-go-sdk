// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstallCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstallCodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstallCodesResponseBody) *DescribeInstallCodesResponse
	GetBody() *DescribeInstallCodesResponseBody
}

type DescribeInstallCodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstallCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstallCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstallCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstallCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstallCodesResponse) GetBody() *DescribeInstallCodesResponseBody {
	return s.Body
}

func (s *DescribeInstallCodesResponse) SetHeaders(v map[string]*string) *DescribeInstallCodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstallCodesResponse) SetStatusCode(v int32) *DescribeInstallCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstallCodesResponse) SetBody(v *DescribeInstallCodesResponseBody) *DescribeInstallCodesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstallCodesResponse) Validate() error {
	return dara.Validate(s)
}
