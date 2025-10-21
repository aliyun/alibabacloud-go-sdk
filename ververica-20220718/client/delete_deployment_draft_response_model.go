// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeploymentDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeploymentDraftResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeploymentDraftResponseBody) *DeleteDeploymentDraftResponse
	GetBody() *DeleteDeploymentDraftResponseBody
}

type DeleteDeploymentDraftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeploymentDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeploymentDraftResponse) GetBody() *DeleteDeploymentDraftResponseBody {
	return s.Body
}

func (s *DeleteDeploymentDraftResponse) SetHeaders(v map[string]*string) *DeleteDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentDraftResponse) SetStatusCode(v int32) *DeleteDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentDraftResponse) SetBody(v *DeleteDeploymentDraftResponseBody) *DeleteDeploymentDraftResponse {
	s.Body = v
	return s
}

func (s *DeleteDeploymentDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
