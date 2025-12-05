// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentsByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentsByNameResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentsByNameResponseBody) *GetDeploymentsByNameResponse
	GetBody() *GetDeploymentsByNameResponseBody
}

type GetDeploymentsByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentsByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentsByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByNameResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentsByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentsByNameResponse) GetBody() *GetDeploymentsByNameResponseBody {
	return s.Body
}

func (s *GetDeploymentsByNameResponse) SetHeaders(v map[string]*string) *GetDeploymentsByNameResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentsByNameResponse) SetStatusCode(v int32) *GetDeploymentsByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentsByNameResponse) SetBody(v *GetDeploymentsByNameResponseBody) *GetDeploymentsByNameResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentsByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
