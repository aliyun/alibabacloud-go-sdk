// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDefaultWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDefaultWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDefaultWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDefaultWorkspaceResponseBody) *UpdateDefaultWorkspaceResponse
	GetBody() *UpdateDefaultWorkspaceResponseBody
}

type UpdateDefaultWorkspaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDefaultWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDefaultWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDefaultWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDefaultWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDefaultWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDefaultWorkspaceResponse) GetBody() *UpdateDefaultWorkspaceResponseBody {
	return s.Body
}

func (s *UpdateDefaultWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateDefaultWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDefaultWorkspaceResponse) SetStatusCode(v int32) *UpdateDefaultWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDefaultWorkspaceResponse) SetBody(v *UpdateDefaultWorkspaceResponseBody) *UpdateDefaultWorkspaceResponse {
	s.Body = v
	return s
}

func (s *UpdateDefaultWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
