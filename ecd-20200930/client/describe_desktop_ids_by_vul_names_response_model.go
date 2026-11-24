// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopIdsByVulNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopIdsByVulNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopIdsByVulNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopIdsByVulNamesResponseBody) *DescribeDesktopIdsByVulNamesResponse
	GetBody() *DescribeDesktopIdsByVulNamesResponseBody
}

type DescribeDesktopIdsByVulNamesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopIdsByVulNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopIdsByVulNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopIdsByVulNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopIdsByVulNamesResponse) GetBody() *DescribeDesktopIdsByVulNamesResponseBody {
	return s.Body
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetHeaders(v map[string]*string) *DescribeDesktopIdsByVulNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetStatusCode(v int32) *DescribeDesktopIdsByVulNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponse) SetBody(v *DescribeDesktopIdsByVulNamesResponseBody) *DescribeDesktopIdsByVulNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
