// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *GetAppDeploymentResponseBody) *GetAppDeploymentResponse
	GetBody() *GetAppDeploymentResponseBody
}

type GetAppDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetAppDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppDeploymentResponse) GetBody() *GetAppDeploymentResponseBody {
	return s.Body
}

func (s *GetAppDeploymentResponse) SetHeaders(v map[string]*string) *GetAppDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetAppDeploymentResponse) SetStatusCode(v int32) *GetAppDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppDeploymentResponse) SetBody(v *GetAppDeploymentResponseBody) *GetAppDeploymentResponse {
	s.Body = v
	return s
}

func (s *GetAppDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
