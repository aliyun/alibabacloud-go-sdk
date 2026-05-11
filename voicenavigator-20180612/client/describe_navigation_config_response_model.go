// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNavigationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNavigationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNavigationConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNavigationConfigResponseBody) *DescribeNavigationConfigResponse
	GetBody() *DescribeNavigationConfigResponseBody
}

type DescribeNavigationConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNavigationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNavigationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNavigationConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeNavigationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNavigationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNavigationConfigResponse) GetBody() *DescribeNavigationConfigResponseBody {
	return s.Body
}

func (s *DescribeNavigationConfigResponse) SetHeaders(v map[string]*string) *DescribeNavigationConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeNavigationConfigResponse) SetStatusCode(v int32) *DescribeNavigationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNavigationConfigResponse) SetBody(v *DescribeNavigationConfigResponseBody) *DescribeNavigationConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeNavigationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
