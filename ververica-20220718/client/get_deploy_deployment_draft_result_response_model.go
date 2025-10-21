// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeployDeploymentDraftResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeployDeploymentDraftResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeployDeploymentDraftResultResponse
	GetStatusCode() *int32
	SetBody(v *GetDeployDeploymentDraftResultResponseBody) *GetDeployDeploymentDraftResultResponse
	GetBody() *GetDeployDeploymentDraftResultResponseBody
}

type GetDeployDeploymentDraftResultResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeployDeploymentDraftResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeployDeploymentDraftResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeployDeploymentDraftResultResponse) GoString() string {
	return s.String()
}

func (s *GetDeployDeploymentDraftResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeployDeploymentDraftResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeployDeploymentDraftResultResponse) GetBody() *GetDeployDeploymentDraftResultResponseBody {
	return s.Body
}

func (s *GetDeployDeploymentDraftResultResponse) SetHeaders(v map[string]*string) *GetDeployDeploymentDraftResultResponse {
	s.Headers = v
	return s
}

func (s *GetDeployDeploymentDraftResultResponse) SetStatusCode(v int32) *GetDeployDeploymentDraftResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeployDeploymentDraftResultResponse) SetBody(v *GetDeployDeploymentDraftResultResponseBody) *GetDeployDeploymentDraftResultResponse {
	s.Body = v
	return s
}

func (s *GetDeployDeploymentDraftResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
