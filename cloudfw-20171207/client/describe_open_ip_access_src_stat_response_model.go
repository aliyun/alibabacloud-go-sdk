// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenIpAccessSrcStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenIpAccessSrcStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenIpAccessSrcStatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenIpAccessSrcStatResponseBody) *DescribeOpenIpAccessSrcStatResponse
	GetBody() *DescribeOpenIpAccessSrcStatResponseBody
}

type DescribeOpenIpAccessSrcStatResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenIpAccessSrcStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenIpAccessSrcStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenIpAccessSrcStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenIpAccessSrcStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenIpAccessSrcStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenIpAccessSrcStatResponse) GetBody() *DescribeOpenIpAccessSrcStatResponseBody {
	return s.Body
}

func (s *DescribeOpenIpAccessSrcStatResponse) SetHeaders(v map[string]*string) *DescribeOpenIpAccessSrcStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponse) SetStatusCode(v int32) *DescribeOpenIpAccessSrcStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponse) SetBody(v *DescribeOpenIpAccessSrcStatResponseBody) *DescribeOpenIpAccessSrcStatResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenIpAccessSrcStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
