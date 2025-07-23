// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentJobResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentJobResponseBody) *ListDeploymentJobResponse
	GetBody() *ListDeploymentJobResponseBody
}

type ListDeploymentJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentJobResponse) GetBody() *ListDeploymentJobResponseBody {
	return s.Body
}

func (s *ListDeploymentJobResponse) SetHeaders(v map[string]*string) *ListDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentJobResponse) SetStatusCode(v int32) *ListDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentJobResponse) SetBody(v *ListDeploymentJobResponseBody) *ListDeploymentJobResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentJobResponse) Validate() error {
	return dara.Validate(s)
}
