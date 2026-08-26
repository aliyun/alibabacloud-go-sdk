// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataAgentThemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataAgentThemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataAgentThemeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataAgentThemeResponseBody) *DescribeDataAgentThemeResponse
	GetBody() *DescribeDataAgentThemeResponseBody
}

type DescribeDataAgentThemeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataAgentThemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataAgentThemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataAgentThemeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataAgentThemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataAgentThemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataAgentThemeResponse) GetBody() *DescribeDataAgentThemeResponseBody {
	return s.Body
}

func (s *DescribeDataAgentThemeResponse) SetHeaders(v map[string]*string) *DescribeDataAgentThemeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataAgentThemeResponse) SetStatusCode(v int32) *DescribeDataAgentThemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataAgentThemeResponse) SetBody(v *DescribeDataAgentThemeResponseBody) *DescribeDataAgentThemeResponse {
	s.Body = v
	return s
}

func (s *DescribeDataAgentThemeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
