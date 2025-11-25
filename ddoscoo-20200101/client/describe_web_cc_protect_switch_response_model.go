// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCcProtectSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebCcProtectSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebCcProtectSwitchResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebCcProtectSwitchResponseBody) *DescribeWebCcProtectSwitchResponse
	GetBody() *DescribeWebCcProtectSwitchResponseBody
}

type DescribeWebCcProtectSwitchResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebCcProtectSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebCcProtectSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCcProtectSwitchResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCcProtectSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebCcProtectSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebCcProtectSwitchResponse) GetBody() *DescribeWebCcProtectSwitchResponseBody {
	return s.Body
}

func (s *DescribeWebCcProtectSwitchResponse) SetHeaders(v map[string]*string) *DescribeWebCcProtectSwitchResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCcProtectSwitchResponse) SetStatusCode(v int32) *DescribeWebCcProtectSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebCcProtectSwitchResponse) SetBody(v *DescribeWebCcProtectSwitchResponseBody) *DescribeWebCcProtectSwitchResponse {
	s.Body = v
	return s
}

func (s *DescribeWebCcProtectSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
