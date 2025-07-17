// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentPackageResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentPackageResponseBody) *GetDeploymentPackageResponse
	GetBody() *GetDeploymentPackageResponseBody
}

type GetDeploymentPackageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentPackageResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentPackageResponse) GetBody() *GetDeploymentPackageResponseBody {
	return s.Body
}

func (s *GetDeploymentPackageResponse) SetHeaders(v map[string]*string) *GetDeploymentPackageResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentPackageResponse) SetStatusCode(v int32) *GetDeploymentPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentPackageResponse) SetBody(v *GetDeploymentPackageResponseBody) *GetDeploymentPackageResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentPackageResponse) Validate() error {
	return dara.Validate(s)
}
