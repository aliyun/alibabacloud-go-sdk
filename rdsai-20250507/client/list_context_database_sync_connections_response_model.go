// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncConnectionsResponseBody) *ListContextDatabaseSyncConnectionsResponse
	GetBody() *ListContextDatabaseSyncConnectionsResponseBody
}

type ListContextDatabaseSyncConnectionsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncConnectionsResponse) GetBody() *ListContextDatabaseSyncConnectionsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncConnectionsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponse) SetBody(v *ListContextDatabaseSyncConnectionsResponseBody) *ListContextDatabaseSyncConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
