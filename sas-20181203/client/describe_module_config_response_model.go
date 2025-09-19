// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModuleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModuleConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModuleConfigResponseBody) *DescribeModuleConfigResponse
	GetBody() *DescribeModuleConfigResponseBody
}

type DescribeModuleConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModuleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModuleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeModuleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModuleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModuleConfigResponse) GetBody() *DescribeModuleConfigResponseBody {
	return s.Body
}

func (s *DescribeModuleConfigResponse) SetHeaders(v map[string]*string) *DescribeModuleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeModuleConfigResponse) SetStatusCode(v int32) *DescribeModuleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModuleConfigResponse) SetBody(v *DescribeModuleConfigResponseBody) *DescribeModuleConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeModuleConfigResponse) Validate() error {
	return dara.Validate(s)
}
