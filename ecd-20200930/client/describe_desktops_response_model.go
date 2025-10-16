// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopsResponseBody) *DescribeDesktopsResponse
	GetBody() *DescribeDesktopsResponseBody
}

type DescribeDesktopsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopsResponse) GetBody() *DescribeDesktopsResponseBody {
	return s.Body
}

func (s *DescribeDesktopsResponse) SetHeaders(v map[string]*string) *DescribeDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsResponse) SetStatusCode(v int32) *DescribeDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopsResponse) SetBody(v *DescribeDesktopsResponseBody) *DescribeDesktopsResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
