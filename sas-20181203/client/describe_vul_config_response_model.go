// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulConfigResponseBody) *DescribeVulConfigResponse
	GetBody() *DescribeVulConfigResponseBody
}

type DescribeVulConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulConfigResponse) GetBody() *DescribeVulConfigResponseBody {
	return s.Body
}

func (s *DescribeVulConfigResponse) SetHeaders(v map[string]*string) *DescribeVulConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulConfigResponse) SetStatusCode(v int32) *DescribeVulConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulConfigResponse) SetBody(v *DescribeVulConfigResponseBody) *DescribeVulConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVulConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
