// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSwitchAzoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSwitchAzoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSwitchAzoneResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSwitchAzoneResponseBody) *DescribeInstanceSwitchAzoneResponse
	GetBody() *DescribeInstanceSwitchAzoneResponseBody
}

type DescribeInstanceSwitchAzoneResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSwitchAzoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSwitchAzoneResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSwitchAzoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSwitchAzoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSwitchAzoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSwitchAzoneResponse) GetBody() *DescribeInstanceSwitchAzoneResponseBody {
	return s.Body
}

func (s *DescribeInstanceSwitchAzoneResponse) SetHeaders(v map[string]*string) *DescribeInstanceSwitchAzoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponse) SetStatusCode(v int32) *DescribeInstanceSwitchAzoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponse) SetBody(v *DescribeInstanceSwitchAzoneResponseBody) *DescribeInstanceSwitchAzoneResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSwitchAzoneResponse) Validate() error {
	return dara.Validate(s)
}
