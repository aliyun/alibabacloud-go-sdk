// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentBySelectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentBySelectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentBySelectorResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentBySelectorResponseBody) *GetDeploymentBySelectorResponse
	GetBody() *GetDeploymentBySelectorResponseBody
}

type GetDeploymentBySelectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentBySelectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentBySelectorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentBySelectorResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentBySelectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentBySelectorResponse) GetBody() *GetDeploymentBySelectorResponseBody {
	return s.Body
}

func (s *GetDeploymentBySelectorResponse) SetHeaders(v map[string]*string) *GetDeploymentBySelectorResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentBySelectorResponse) SetStatusCode(v int32) *GetDeploymentBySelectorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentBySelectorResponse) SetBody(v *GetDeploymentBySelectorResponseBody) *GetDeploymentBySelectorResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentBySelectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
