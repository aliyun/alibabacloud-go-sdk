// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopInfoResponseBody) *DescribeDesktopInfoResponse
	GetBody() *DescribeDesktopInfoResponseBody
}

type DescribeDesktopInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopInfoResponse) GetBody() *DescribeDesktopInfoResponseBody {
	return s.Body
}

func (s *DescribeDesktopInfoResponse) SetHeaders(v map[string]*string) *DescribeDesktopInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopInfoResponse) SetStatusCode(v int32) *DescribeDesktopInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopInfoResponse) SetBody(v *DescribeDesktopInfoResponseBody) *DescribeDesktopInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
