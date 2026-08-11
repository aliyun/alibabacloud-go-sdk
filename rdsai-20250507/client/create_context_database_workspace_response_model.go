// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDatabaseWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDatabaseWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDatabaseWorkspaceResponseBody) *CreateContextDatabaseWorkspaceResponse
	GetBody() *CreateContextDatabaseWorkspaceResponseBody
}

type CreateContextDatabaseWorkspaceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDatabaseWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDatabaseWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDatabaseWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDatabaseWorkspaceResponse) GetBody() *CreateContextDatabaseWorkspaceResponseBody {
	return s.Body
}

func (s *CreateContextDatabaseWorkspaceResponse) SetHeaders(v map[string]*string) *CreateContextDatabaseWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponse) SetStatusCode(v int32) *CreateContextDatabaseWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponse) SetBody(v *CreateContextDatabaseWorkspaceResponseBody) *CreateContextDatabaseWorkspaceResponse {
	s.Body = v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
