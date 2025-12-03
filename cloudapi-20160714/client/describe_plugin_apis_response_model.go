// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginApisResponseBody) *DescribePluginApisResponse
	GetBody() *DescribePluginApisResponseBody
}

type DescribePluginApisResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginApisResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginApisResponse) GetBody() *DescribePluginApisResponseBody {
	return s.Body
}

func (s *DescribePluginApisResponse) SetHeaders(v map[string]*string) *DescribePluginApisResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginApisResponse) SetStatusCode(v int32) *DescribePluginApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginApisResponse) SetBody(v *DescribePluginApisResponseBody) *DescribePluginApisResponse {
	s.Body = v
	return s
}

func (s *DescribePluginApisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
