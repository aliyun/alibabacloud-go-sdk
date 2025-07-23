// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeploymentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeploymentJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeploymentJobResponseBody) *DescribeDeploymentJobResponse
	GetBody() *DescribeDeploymentJobResponseBody
}

type DescribeDeploymentJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeploymentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeploymentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeploymentJobResponse) GetBody() *DescribeDeploymentJobResponseBody {
	return s.Body
}

func (s *DescribeDeploymentJobResponse) SetHeaders(v map[string]*string) *DescribeDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeploymentJobResponse) SetStatusCode(v int32) *DescribeDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeploymentJobResponse) SetBody(v *DescribeDeploymentJobResponseBody) *DescribeDeploymentJobResponse {
	s.Body = v
	return s
}

func (s *DescribeDeploymentJobResponse) Validate() error {
	return dara.Validate(s)
}
