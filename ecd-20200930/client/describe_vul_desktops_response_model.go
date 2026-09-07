// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulDesktopsResponseBody) *DescribeVulDesktopsResponse
	GetBody() *DescribeVulDesktopsResponseBody
}

type DescribeVulDesktopsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulDesktopsResponse) GetBody() *DescribeVulDesktopsResponseBody {
	return s.Body
}

func (s *DescribeVulDesktopsResponse) SetHeaders(v map[string]*string) *DescribeVulDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulDesktopsResponse) SetStatusCode(v int32) *DescribeVulDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulDesktopsResponse) SetBody(v *DescribeVulDesktopsResponseBody) *DescribeVulDesktopsResponse {
	s.Body = v
	return s
}

func (s *DescribeVulDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
