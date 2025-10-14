// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnabledCrossRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnabledCrossRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnabledCrossRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnabledCrossRegionsResponseBody) *DescribeEnabledCrossRegionsResponse
	GetBody() *DescribeEnabledCrossRegionsResponseBody
}

type DescribeEnabledCrossRegionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnabledCrossRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnabledCrossRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledCrossRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnabledCrossRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnabledCrossRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnabledCrossRegionsResponse) GetBody() *DescribeEnabledCrossRegionsResponseBody {
	return s.Body
}

func (s *DescribeEnabledCrossRegionsResponse) SetHeaders(v map[string]*string) *DescribeEnabledCrossRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnabledCrossRegionsResponse) SetStatusCode(v int32) *DescribeEnabledCrossRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnabledCrossRegionsResponse) SetBody(v *DescribeEnabledCrossRegionsResponseBody) *DescribeEnabledCrossRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeEnabledCrossRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
