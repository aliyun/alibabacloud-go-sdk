// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlockedRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBlockedRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBlockedRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBlockedRegionsResponseBody) *DescribeBlockedRegionsResponse
	GetBody() *DescribeBlockedRegionsResponseBody
}

type DescribeBlockedRegionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBlockedRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBlockedRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockedRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlockedRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBlockedRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBlockedRegionsResponse) GetBody() *DescribeBlockedRegionsResponseBody {
	return s.Body
}

func (s *DescribeBlockedRegionsResponse) SetHeaders(v map[string]*string) *DescribeBlockedRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlockedRegionsResponse) SetStatusCode(v int32) *DescribeBlockedRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlockedRegionsResponse) SetBody(v *DescribeBlockedRegionsResponseBody) *DescribeBlockedRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeBlockedRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
