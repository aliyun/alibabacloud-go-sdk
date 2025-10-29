// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentPackageFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeploymentPackageFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeploymentPackageFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListDeploymentPackageFilesResponseBody) *ListDeploymentPackageFilesResponse
	GetBody() *ListDeploymentPackageFilesResponseBody
}

type ListDeploymentPackageFilesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentPackageFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentPackageFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentPackageFilesResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentPackageFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeploymentPackageFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeploymentPackageFilesResponse) GetBody() *ListDeploymentPackageFilesResponseBody {
	return s.Body
}

func (s *ListDeploymentPackageFilesResponse) SetHeaders(v map[string]*string) *ListDeploymentPackageFilesResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentPackageFilesResponse) SetStatusCode(v int32) *ListDeploymentPackageFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentPackageFilesResponse) SetBody(v *ListDeploymentPackageFilesResponseBody) *ListDeploymentPackageFilesResponse {
	s.Body = v
	return s
}

func (s *ListDeploymentPackageFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
