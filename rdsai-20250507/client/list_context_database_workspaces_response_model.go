// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseWorkspacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseWorkspacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseWorkspacesResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseWorkspacesResponseBody) *ListContextDatabaseWorkspacesResponse
	GetBody() *ListContextDatabaseWorkspacesResponseBody
}

type ListContextDatabaseWorkspacesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseWorkspacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseWorkspacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseWorkspacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseWorkspacesResponse) GetBody() *ListContextDatabaseWorkspacesResponseBody {
	return s.Body
}

func (s *ListContextDatabaseWorkspacesResponse) SetHeaders(v map[string]*string) *ListContextDatabaseWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseWorkspacesResponse) SetStatusCode(v int32) *ListContextDatabaseWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseWorkspacesResponse) SetBody(v *ListContextDatabaseWorkspacesResponseBody) *ListContextDatabaseWorkspacesResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseWorkspacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
