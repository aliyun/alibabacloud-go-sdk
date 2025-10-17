// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalDesktopsResponseBody) *DescribeGlobalDesktopsResponse
	GetBody() *DescribeGlobalDesktopsResponseBody
}

type DescribeGlobalDesktopsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalDesktopsResponse) GetBody() *DescribeGlobalDesktopsResponseBody {
	return s.Body
}

func (s *DescribeGlobalDesktopsResponse) SetHeaders(v map[string]*string) *DescribeGlobalDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDesktopsResponse) SetStatusCode(v int32) *DescribeGlobalDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDesktopsResponse) SetBody(v *DescribeGlobalDesktopsResponseBody) *DescribeGlobalDesktopsResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
