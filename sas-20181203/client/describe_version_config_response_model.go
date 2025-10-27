// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVersionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVersionConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVersionConfigResponseBody) *DescribeVersionConfigResponse
	GetBody() *DescribeVersionConfigResponseBody
}

type DescribeVersionConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVersionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVersionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVersionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVersionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVersionConfigResponse) GetBody() *DescribeVersionConfigResponseBody {
	return s.Body
}

func (s *DescribeVersionConfigResponse) SetHeaders(v map[string]*string) *DescribeVersionConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVersionConfigResponse) SetStatusCode(v int32) *DescribeVersionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVersionConfigResponse) SetBody(v *DescribeVersionConfigResponseBody) *DescribeVersionConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVersionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
