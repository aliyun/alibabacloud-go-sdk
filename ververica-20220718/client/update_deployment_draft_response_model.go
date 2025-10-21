// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentDraftResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentDraftResponseBody) *UpdateDeploymentDraftResponse
	GetBody() *UpdateDeploymentDraftResponseBody
}

type UpdateDeploymentDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentDraftResponse) GetBody() *UpdateDeploymentDraftResponseBody {
	return s.Body
}

func (s *UpdateDeploymentDraftResponse) SetHeaders(v map[string]*string) *UpdateDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentDraftResponse) SetStatusCode(v int32) *UpdateDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentDraftResponse) SetBody(v *UpdateDeploymentDraftResponseBody) *UpdateDeploymentDraftResponse {
	s.Body = v
	return s
}

func (s *UpdateDeploymentDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
