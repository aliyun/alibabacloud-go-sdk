// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentDraftLockResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentDraftLockResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentDraftLockResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentDraftLockResponseBody) *GetDeploymentDraftLockResponse
	GetBody() *GetDeploymentDraftLockResponseBody
}

type GetDeploymentDraftLockResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentDraftLockResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentDraftLockResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentDraftLockResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentDraftLockResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentDraftLockResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentDraftLockResponse) GetBody() *GetDeploymentDraftLockResponseBody {
	return s.Body
}

func (s *GetDeploymentDraftLockResponse) SetHeaders(v map[string]*string) *GetDeploymentDraftLockResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentDraftLockResponse) SetStatusCode(v int32) *GetDeploymentDraftLockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentDraftLockResponse) SetBody(v *GetDeploymentDraftLockResponseBody) *GetDeploymentDraftLockResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentDraftLockResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
