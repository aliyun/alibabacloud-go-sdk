// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *EnvironmentDeploymentSpec) *PreviewEnvironmentResponse
	GetBody() *EnvironmentDeploymentSpec
}

type PreviewEnvironmentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnvironmentDeploymentSpec `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *PreviewEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewEnvironmentResponse) GetBody() *EnvironmentDeploymentSpec {
	return s.Body
}

func (s *PreviewEnvironmentResponse) SetHeaders(v map[string]*string) *PreviewEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *PreviewEnvironmentResponse) SetStatusCode(v int32) *PreviewEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewEnvironmentResponse) SetBody(v *EnvironmentDeploymentSpec) *PreviewEnvironmentResponse {
	s.Body = v
	return s
}

func (s *PreviewEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
