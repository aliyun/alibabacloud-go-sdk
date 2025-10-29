// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentPackagesResponseBody) *ListDeploymentPackagesResponse
	GetBody() *ListDeploymentPackagesResponseBody
}

type ListDeploymentPackagesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentPackagesResponse) GetBody() *ListDeploymentPackagesResponseBody {
	return s.Body
}

func (s *ListDeploymentPackagesResponse) SetHeaders(v map[string]*string) *ListDeploymentPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentPackagesResponse) SetStatusCode(v int32) *ListDeploymentPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentPackagesResponse) SetBody(v *ListDeploymentPackagesResponseBody) *ListDeploymentPackagesResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
