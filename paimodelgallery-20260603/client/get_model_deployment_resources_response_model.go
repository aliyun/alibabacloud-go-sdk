// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelDeploymentResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModelDeploymentResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModelDeploymentResourcesResponse
	GetStatusCode() *int32
	SetBody(v *GetModelDeploymentResourcesResponseBody) *GetModelDeploymentResourcesResponse
	GetBody() *GetModelDeploymentResourcesResponseBody
}

type GetModelDeploymentResourcesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelDeploymentResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelDeploymentResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModelDeploymentResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModelDeploymentResourcesResponse) GetBody() *GetModelDeploymentResourcesResponseBody {
	return s.Body
}

func (s *GetModelDeploymentResourcesResponse) SetHeaders(v map[string]*string) *GetModelDeploymentResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetModelDeploymentResourcesResponse) SetStatusCode(v int32) *GetModelDeploymentResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelDeploymentResourcesResponse) SetBody(v *GetModelDeploymentResourcesResponseBody) *GetModelDeploymentResourcesResponse {
	s.Body = v
	return s
}

func (s *GetModelDeploymentResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
