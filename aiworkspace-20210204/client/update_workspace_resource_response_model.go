// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWorkspaceResourceResponseBody) *UpdateWorkspaceResourceResponse
	GetBody() *UpdateWorkspaceResourceResponseBody
}

type UpdateWorkspaceResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceResourceResponse) GetBody() *UpdateWorkspaceResourceResponseBody {
	return s.Body
}

func (s *UpdateWorkspaceResourceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResourceResponse) SetStatusCode(v int32) *UpdateWorkspaceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResourceResponse) SetBody(v *UpdateWorkspaceResourceResponseBody) *UpdateWorkspaceResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceResourceResponse) Validate() error {
	return dara.Validate(s)
}
