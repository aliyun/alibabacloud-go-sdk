// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssStatusResponseBody) *DescribeOssStatusResponse
	GetBody() *DescribeOssStatusResponseBody
}

type DescribeOssStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssStatusResponse) GetBody() *DescribeOssStatusResponseBody {
	return s.Body
}

func (s *DescribeOssStatusResponse) SetHeaders(v map[string]*string) *DescribeOssStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssStatusResponse) SetStatusCode(v int32) *DescribeOssStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssStatusResponse) SetBody(v *DescribeOssStatusResponseBody) *DescribeOssStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeOssStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
