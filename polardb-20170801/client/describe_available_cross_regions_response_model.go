// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableCrossRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableCrossRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableCrossRegionsResponseBody) *DescribeAvailableCrossRegionsResponse
	GetBody() *DescribeAvailableCrossRegionsResponseBody
}

type DescribeAvailableCrossRegionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableCrossRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableCrossRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableCrossRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableCrossRegionsResponse) GetBody() *DescribeAvailableCrossRegionsResponseBody {
	return s.Body
}

func (s *DescribeAvailableCrossRegionsResponse) SetHeaders(v map[string]*string) *DescribeAvailableCrossRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableCrossRegionsResponse) SetStatusCode(v int32) *DescribeAvailableCrossRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableCrossRegionsResponse) SetBody(v *DescribeAvailableCrossRegionsResponseBody) *DescribeAvailableCrossRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableCrossRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
