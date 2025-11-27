// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRenderingInstanceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRenderingInstanceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRenderingInstanceSettingsResponseBody) *DescribeRenderingInstanceSettingsResponse
	GetBody() *DescribeRenderingInstanceSettingsResponseBody
}

type DescribeRenderingInstanceSettingsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRenderingInstanceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRenderingInstanceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceSettingsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRenderingInstanceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRenderingInstanceSettingsResponse) GetBody() *DescribeRenderingInstanceSettingsResponseBody {
	return s.Body
}

func (s *DescribeRenderingInstanceSettingsResponse) SetHeaders(v map[string]*string) *DescribeRenderingInstanceSettingsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponse) SetStatusCode(v int32) *DescribeRenderingInstanceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponse) SetBody(v *DescribeRenderingInstanceSettingsResponseBody) *DescribeRenderingInstanceSettingsResponse {
	s.Body = v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
