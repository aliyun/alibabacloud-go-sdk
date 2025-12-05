// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetValidateDeploymentDraftResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetValidateDeploymentDraftResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetValidateDeploymentDraftResultResponse
	GetStatusCode() *int32
	SetBody(v *GetValidateDeploymentDraftResultResponseBody) *GetValidateDeploymentDraftResultResponse
	GetBody() *GetValidateDeploymentDraftResultResponseBody
}

type GetValidateDeploymentDraftResultResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetValidateDeploymentDraftResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetValidateDeploymentDraftResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetValidateDeploymentDraftResultResponse) GoString() string {
	return s.String()
}

func (s *GetValidateDeploymentDraftResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetValidateDeploymentDraftResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetValidateDeploymentDraftResultResponse) GetBody() *GetValidateDeploymentDraftResultResponseBody {
	return s.Body
}

func (s *GetValidateDeploymentDraftResultResponse) SetHeaders(v map[string]*string) *GetValidateDeploymentDraftResultResponse {
	s.Headers = v
	return s
}

func (s *GetValidateDeploymentDraftResultResponse) SetStatusCode(v int32) *GetValidateDeploymentDraftResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetValidateDeploymentDraftResultResponse) SetBody(v *GetValidateDeploymentDraftResultResponseBody) *GetValidateDeploymentDraftResultResponse {
	s.Body = v
	return s
}

func (s *GetValidateDeploymentDraftResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
