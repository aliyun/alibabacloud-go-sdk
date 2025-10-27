// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageFixCycleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageFixCycleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageFixCycleConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageFixCycleConfigResponseBody) *DescribeImageFixCycleConfigResponse
	GetBody() *DescribeImageFixCycleConfigResponseBody
}

type DescribeImageFixCycleConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageFixCycleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageFixCycleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageFixCycleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageFixCycleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageFixCycleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageFixCycleConfigResponse) GetBody() *DescribeImageFixCycleConfigResponseBody {
	return s.Body
}

func (s *DescribeImageFixCycleConfigResponse) SetHeaders(v map[string]*string) *DescribeImageFixCycleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageFixCycleConfigResponse) SetStatusCode(v int32) *DescribeImageFixCycleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageFixCycleConfigResponse) SetBody(v *DescribeImageFixCycleConfigResponseBody) *DescribeImageFixCycleConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeImageFixCycleConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
