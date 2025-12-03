// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginsResponseBody) *DescribePluginsResponse
	GetBody() *DescribePluginsResponseBody
}

type DescribePluginsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginsResponse) GetBody() *DescribePluginsResponseBody {
	return s.Body
}

func (s *DescribePluginsResponse) SetHeaders(v map[string]*string) *DescribePluginsResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginsResponse) SetStatusCode(v int32) *DescribePluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginsResponse) SetBody(v *DescribePluginsResponseBody) *DescribePluginsResponse {
	s.Body = v
	return s
}

func (s *DescribePluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
