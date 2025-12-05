// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentsByLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentsByLabelResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentsByLabelResponseBody) *GetDeploymentsByLabelResponse
	GetBody() *GetDeploymentsByLabelResponseBody
}

type GetDeploymentsByLabelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentsByLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentsByLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByLabelResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentsByLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentsByLabelResponse) GetBody() *GetDeploymentsByLabelResponseBody {
	return s.Body
}

func (s *GetDeploymentsByLabelResponse) SetHeaders(v map[string]*string) *GetDeploymentsByLabelResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentsByLabelResponse) SetStatusCode(v int32) *GetDeploymentsByLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentsByLabelResponse) SetBody(v *GetDeploymentsByLabelResponseBody) *GetDeploymentsByLabelResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentsByLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
