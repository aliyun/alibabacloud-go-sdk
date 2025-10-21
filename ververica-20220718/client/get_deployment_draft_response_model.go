// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentDraftResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentDraftResponseBody) *GetDeploymentDraftResponse
	GetBody() *GetDeploymentDraftResponseBody
}

type GetDeploymentDraftResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentDraftResponse) GetBody() *GetDeploymentDraftResponseBody {
	return s.Body
}

func (s *GetDeploymentDraftResponse) SetHeaders(v map[string]*string) *GetDeploymentDraftResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentDraftResponse) SetStatusCode(v int32) *GetDeploymentDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentDraftResponse) SetBody(v *GetDeploymentDraftResponseBody) *GetDeploymentDraftResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
