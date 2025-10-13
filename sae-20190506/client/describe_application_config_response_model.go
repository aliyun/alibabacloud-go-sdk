// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationConfigResponseBody) *DescribeApplicationConfigResponse
	GetBody() *DescribeApplicationConfigResponseBody
}

type DescribeApplicationConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationConfigResponse) GetBody() *DescribeApplicationConfigResponseBody {
	return s.Body
}

func (s *DescribeApplicationConfigResponse) SetHeaders(v map[string]*string) *DescribeApplicationConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationConfigResponse) SetStatusCode(v int32) *DescribeApplicationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationConfigResponse) SetBody(v *DescribeApplicationConfigResponseBody) *DescribeApplicationConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
