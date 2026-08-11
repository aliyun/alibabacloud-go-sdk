// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseWorkspaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextDatabaseWorkspaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextDatabaseWorkspaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextDatabaseWorkspaceResponseBody) *DeleteContextDatabaseWorkspaceResponse
	GetBody() *DeleteContextDatabaseWorkspaceResponseBody
}

type DeleteContextDatabaseWorkspaceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextDatabaseWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextDatabaseWorkspaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseWorkspaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextDatabaseWorkspaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextDatabaseWorkspaceResponse) GetBody() *DeleteContextDatabaseWorkspaceResponseBody {
	return s.Body
}

func (s *DeleteContextDatabaseWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteContextDatabaseWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponse) SetStatusCode(v int32) *DeleteContextDatabaseWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponse) SetBody(v *DeleteContextDatabaseWorkspaceResponseBody) *DeleteContextDatabaseWorkspaceResponse {
	s.Body = v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
