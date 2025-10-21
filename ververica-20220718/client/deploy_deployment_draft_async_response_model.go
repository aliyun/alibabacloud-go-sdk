// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployDeploymentDraftAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployDeploymentDraftAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployDeploymentDraftAsyncResponse
	GetStatusCode() *int32
	SetBody(v *DeployDeploymentDraftAsyncResponseBody) *DeployDeploymentDraftAsyncResponse
	GetBody() *DeployDeploymentDraftAsyncResponseBody
}

type DeployDeploymentDraftAsyncResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployDeploymentDraftAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployDeploymentDraftAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployDeploymentDraftAsyncResponse) GoString() string {
	return s.String()
}

func (s *DeployDeploymentDraftAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployDeploymentDraftAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployDeploymentDraftAsyncResponse) GetBody() *DeployDeploymentDraftAsyncResponseBody {
	return s.Body
}

func (s *DeployDeploymentDraftAsyncResponse) SetHeaders(v map[string]*string) *DeployDeploymentDraftAsyncResponse {
	s.Headers = v
	return s
}

func (s *DeployDeploymentDraftAsyncResponse) SetStatusCode(v int32) *DeployDeploymentDraftAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployDeploymentDraftAsyncResponse) SetBody(v *DeployDeploymentDraftAsyncResponseBody) *DeployDeploymentDraftAsyncResponse {
	s.Body = v
	return s
}

func (s *DeployDeploymentDraftAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
