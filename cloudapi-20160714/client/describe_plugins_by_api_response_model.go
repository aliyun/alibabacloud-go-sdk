// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsByApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePluginsByApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePluginsByApiResponse
	GetStatusCode() *int32
	SetBody(v *DescribePluginsByApiResponseBody) *DescribePluginsByApiResponse
	GetBody() *DescribePluginsByApiResponseBody
}

type DescribePluginsByApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePluginsByApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePluginsByApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePluginsByApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePluginsByApiResponse) GetBody() *DescribePluginsByApiResponseBody {
	return s.Body
}

func (s *DescribePluginsByApiResponse) SetHeaders(v map[string]*string) *DescribePluginsByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginsByApiResponse) SetStatusCode(v int32) *DescribePluginsByApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePluginsByApiResponse) SetBody(v *DescribePluginsByApiResponseBody) *DescribePluginsByApiResponse {
	s.Body = v
	return s
}

func (s *DescribePluginsByApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
