// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvisorResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvisorResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvisorResourcesResponseBody) *DescribeAdvisorResourcesResponse
	GetBody() *DescribeAdvisorResourcesResponseBody
}

type DescribeAdvisorResourcesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvisorResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvisorResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvisorResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvisorResourcesResponse) GetBody() *DescribeAdvisorResourcesResponseBody {
	return s.Body
}

func (s *DescribeAdvisorResourcesResponse) SetHeaders(v map[string]*string) *DescribeAdvisorResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvisorResourcesResponse) SetStatusCode(v int32) *DescribeAdvisorResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvisorResourcesResponse) SetBody(v *DescribeAdvisorResourcesResponseBody) *DescribeAdvisorResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvisorResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
