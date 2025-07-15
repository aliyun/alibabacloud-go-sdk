// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployedApisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeployedApisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeployedApisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeployedApisResponseBody) *DescribeDeployedApisResponse
	GetBody() *DescribeDeployedApisResponseBody
}

type DescribeDeployedApisResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeployedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeployedApisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeployedApisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeployedApisResponse) GetBody() *DescribeDeployedApisResponseBody {
	return s.Body
}

func (s *DescribeDeployedApisResponse) SetHeaders(v map[string]*string) *DescribeDeployedApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployedApisResponse) SetStatusCode(v int32) *DescribeDeployedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeployedApisResponse) SetBody(v *DescribeDeployedApisResponseBody) *DescribeDeployedApisResponse {
	s.Body = v
	return s
}

func (s *DescribeDeployedApisResponse) Validate() error {
	return dara.Validate(s)
}
