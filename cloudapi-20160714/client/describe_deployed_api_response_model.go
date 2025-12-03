// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployedApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeployedApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeployedApiResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeployedApiResponseBody) *DescribeDeployedApiResponse
	GetBody() *DescribeDeployedApiResponseBody
}

type DescribeDeployedApiResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeployedApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeployedApiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeployedApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeployedApiResponse) GetBody() *DescribeDeployedApiResponseBody {
	return s.Body
}

func (s *DescribeDeployedApiResponse) SetHeaders(v map[string]*string) *DescribeDeployedApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployedApiResponse) SetStatusCode(v int32) *DescribeDeployedApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeployedApiResponse) SetBody(v *DescribeDeployedApiResponseBody) *DescribeDeployedApiResponse {
	s.Body = v
	return s
}

func (s *DescribeDeployedApiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
