// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWorkspaceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWorkspaceResourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateWorkspaceResourceResponseBody) *CreateWorkspaceResourceResponse
	GetBody() *CreateWorkspaceResourceResponseBody
}

type CreateWorkspaceResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWorkspaceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWorkspaceResourceResponse) GetBody() *CreateWorkspaceResourceResponseBody {
	return s.Body
}

func (s *CreateWorkspaceResourceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResourceResponse) SetStatusCode(v int32) *CreateWorkspaceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResourceResponse) SetBody(v *CreateWorkspaceResourceResponseBody) *CreateWorkspaceResourceResponse {
	s.Body = v
	return s
}

func (s *CreateWorkspaceResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
