// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginTemplatesResponseBody) *DescribePluginTemplatesResponse
	GetBody() *DescribePluginTemplatesResponseBody
}

type DescribePluginTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginTemplatesResponse) GetBody() *DescribePluginTemplatesResponseBody {
	return s.Body
}

func (s *DescribePluginTemplatesResponse) SetHeaders(v map[string]*string) *DescribePluginTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginTemplatesResponse) SetStatusCode(v int32) *DescribePluginTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginTemplatesResponse) SetBody(v *DescribePluginTemplatesResponseBody) *DescribePluginTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribePluginTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
