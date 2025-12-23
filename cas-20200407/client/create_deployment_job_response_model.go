// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeploymentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeploymentJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeploymentJobResponseBody) *CreateDeploymentJobResponse
	GetBody() *CreateDeploymentJobResponseBody
}

type CreateDeploymentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeploymentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeploymentJobResponse) GetBody() *CreateDeploymentJobResponseBody {
	return s.Body
}

func (s *CreateDeploymentJobResponse) SetHeaders(v map[string]*string) *CreateDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentJobResponse) SetStatusCode(v int32) *CreateDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentJobResponse) SetBody(v *CreateDeploymentJobResponseBody) *CreateDeploymentJobResponse {
	s.Body = v
	return s
}

func (s *CreateDeploymentJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
