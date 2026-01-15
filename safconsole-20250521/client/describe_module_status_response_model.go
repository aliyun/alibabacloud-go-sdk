// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModuleStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModuleStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModuleStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModuleStatusResponseBody) *DescribeModuleStatusResponse
	GetBody() *DescribeModuleStatusResponseBody
}

type DescribeModuleStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModuleStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModuleStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModuleStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeModuleStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModuleStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModuleStatusResponse) GetBody() *DescribeModuleStatusResponseBody {
	return s.Body
}

func (s *DescribeModuleStatusResponse) SetHeaders(v map[string]*string) *DescribeModuleStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeModuleStatusResponse) SetStatusCode(v int32) *DescribeModuleStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModuleStatusResponse) SetBody(v *DescribeModuleStatusResponseBody) *DescribeModuleStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeModuleStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
