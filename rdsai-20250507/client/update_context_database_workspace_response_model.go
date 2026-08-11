// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseWorkspaceResponseBody) *UpdateContextDatabaseWorkspaceResponse
	GetBody() *UpdateContextDatabaseWorkspaceResponseBody
}

type UpdateContextDatabaseWorkspaceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseWorkspaceResponse) GetBody() *UpdateContextDatabaseWorkspaceResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponse) SetStatusCode(v int32) *UpdateContextDatabaseWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponse) SetBody(v *UpdateContextDatabaseWorkspaceResponseBody) *UpdateContextDatabaseWorkspaceResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
