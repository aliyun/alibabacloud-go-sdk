// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEipResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEipResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEipResourcesResponseBody) *DescribeEipResourcesResponse
	GetBody() *DescribeEipResourcesResponseBody
}

type DescribeEipResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEipResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEipResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEipResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEipResourcesResponse) GetBody() *DescribeEipResourcesResponseBody {
	return s.Body
}

func (s *DescribeEipResourcesResponse) SetHeaders(v map[string]*string) *DescribeEipResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipResourcesResponse) SetStatusCode(v int32) *DescribeEipResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEipResourcesResponse) SetBody(v *DescribeEipResourcesResponseBody) *DescribeEipResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeEipResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
