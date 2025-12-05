// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDeploymentDraftAsyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateDeploymentDraftAsyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateDeploymentDraftAsyncResponse
	GetStatusCode() *int32
	SetBody(v *ValidateDeploymentDraftAsyncResponseBody) *ValidateDeploymentDraftAsyncResponse
	GetBody() *ValidateDeploymentDraftAsyncResponseBody
}

type ValidateDeploymentDraftAsyncResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateDeploymentDraftAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDeploymentDraftAsyncResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateDeploymentDraftAsyncResponse) GoString() string {
	return s.String()
}

func (s *ValidateDeploymentDraftAsyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateDeploymentDraftAsyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateDeploymentDraftAsyncResponse) GetBody() *ValidateDeploymentDraftAsyncResponseBody {
	return s.Body
}

func (s *ValidateDeploymentDraftAsyncResponse) SetHeaders(v map[string]*string) *ValidateDeploymentDraftAsyncResponse {
	s.Headers = v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponse) SetStatusCode(v int32) *ValidateDeploymentDraftAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponse) SetBody(v *ValidateDeploymentDraftAsyncResponseBody) *ValidateDeploymentDraftAsyncResponse {
	s.Body = v
	return s
}

func (s *ValidateDeploymentDraftAsyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
