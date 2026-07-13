// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopRiskyResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTopRiskyResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTopRiskyResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTopRiskyResourcesResponseBody) *DescribeTopRiskyResourcesResponse
	GetBody() *DescribeTopRiskyResourcesResponseBody
}

type DescribeTopRiskyResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTopRiskyResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTopRiskyResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTopRiskyResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTopRiskyResourcesResponse) GetBody() *DescribeTopRiskyResourcesResponseBody {
	return s.Body
}

func (s *DescribeTopRiskyResourcesResponse) SetHeaders(v map[string]*string) *DescribeTopRiskyResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopRiskyResourcesResponse) SetStatusCode(v int32) *DescribeTopRiskyResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponse) SetBody(v *DescribeTopRiskyResourcesResponseBody) *DescribeTopRiskyResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeTopRiskyResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
