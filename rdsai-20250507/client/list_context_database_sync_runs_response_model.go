// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncRunsResponseBody) *ListContextDatabaseSyncRunsResponse
	GetBody() *ListContextDatabaseSyncRunsResponseBody
}

type ListContextDatabaseSyncRunsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncRunsResponse) GetBody() *ListContextDatabaseSyncRunsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncRunsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncRunsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncRunsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncRunsResponse) SetBody(v *ListContextDatabaseSyncRunsResponseBody) *ListContextDatabaseSyncRunsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
