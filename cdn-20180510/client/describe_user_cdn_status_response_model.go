// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserCdnStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserCdnStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserCdnStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserCdnStatusResponseBody) *DescribeUserCdnStatusResponse
	GetBody() *DescribeUserCdnStatusResponseBody
}

type DescribeUserCdnStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserCdnStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserCdnStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserCdnStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserCdnStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserCdnStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserCdnStatusResponse) GetBody() *DescribeUserCdnStatusResponseBody {
	return s.Body
}

func (s *DescribeUserCdnStatusResponse) SetHeaders(v map[string]*string) *DescribeUserCdnStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserCdnStatusResponse) SetStatusCode(v int32) *DescribeUserCdnStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserCdnStatusResponse) SetBody(v *DescribeUserCdnStatusResponseBody) *DescribeUserCdnStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserCdnStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
