// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelDeploymentSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelDeploymentSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelDeploymentSpecResponse
	GetStatusCode() *int32
	SetBody(v *GetModelDeploymentSpecResponseBody) *GetModelDeploymentSpecResponse
	GetBody() *GetModelDeploymentSpecResponseBody
}

type GetModelDeploymentSpecResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelDeploymentSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelDeploymentSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentSpecResponse) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelDeploymentSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelDeploymentSpecResponse) GetBody() *GetModelDeploymentSpecResponseBody {
	return s.Body
}

func (s *GetModelDeploymentSpecResponse) SetHeaders(v map[string]*string) *GetModelDeploymentSpecResponse {
	s.Headers = v
	return s
}

func (s *GetModelDeploymentSpecResponse) SetStatusCode(v int32) *GetModelDeploymentSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelDeploymentSpecResponse) SetBody(v *GetModelDeploymentSpecResponseBody) *GetModelDeploymentSpecResponse {
	s.Body = v
	return s
}

func (s *GetModelDeploymentSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
