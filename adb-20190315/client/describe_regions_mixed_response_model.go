// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsMixedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionsMixedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionsMixedResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionsMixedResponseBody) *DescribeRegionsMixedResponse
	GetBody() *DescribeRegionsMixedResponseBody
}

type DescribeRegionsMixedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsMixedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsMixedResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsMixedResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsMixedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionsMixedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionsMixedResponse) GetBody() *DescribeRegionsMixedResponseBody {
	return s.Body
}

func (s *DescribeRegionsMixedResponse) SetHeaders(v map[string]*string) *DescribeRegionsMixedResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsMixedResponse) SetStatusCode(v int32) *DescribeRegionsMixedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsMixedResponse) SetBody(v *DescribeRegionsMixedResponseBody) *DescribeRegionsMixedResponse {
	s.Body = v
	return s
}

func (s *DescribeRegionsMixedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
