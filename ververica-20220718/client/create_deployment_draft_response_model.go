// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeploymentDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeploymentDraftResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeploymentDraftResponseBody) *CreateDeploymentDraftResponse
	GetBody() *CreateDeploymentDraftResponseBody
}

type CreateDeploymentDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeploymentDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeploymentDraftResponse) GetBody() *CreateDeploymentDraftResponseBody {
	return s.Body
}

func (s *CreateDeploymentDraftResponse) SetHeaders(v map[string]*string) *CreateDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentDraftResponse) SetStatusCode(v int32) *CreateDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentDraftResponse) SetBody(v *CreateDeploymentDraftResponseBody) *CreateDeploymentDraftResponse {
	s.Body = v
	return s
}

func (s *CreateDeploymentDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
