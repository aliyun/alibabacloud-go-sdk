// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployApiTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeployApiTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeployApiTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeployApiTaskResponseBody) *DescribeDeployApiTaskResponse
	GetBody() *DescribeDeployApiTaskResponseBody
}

type DescribeDeployApiTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeployApiTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeployApiTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployApiTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeployApiTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeployApiTaskResponse) GetBody() *DescribeDeployApiTaskResponseBody {
	return s.Body
}

func (s *DescribeDeployApiTaskResponse) SetHeaders(v map[string]*string) *DescribeDeployApiTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployApiTaskResponse) SetStatusCode(v int32) *DescribeDeployApiTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeployApiTaskResponse) SetBody(v *DescribeDeployApiTaskResponseBody) *DescribeDeployApiTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeDeployApiTaskResponse) Validate() error {
	return dara.Validate(s)
}
