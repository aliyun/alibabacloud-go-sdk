// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncRunLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncRunLogsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncRunLogsResponseBody) *ListContextDatabaseSyncRunLogsResponse
	GetBody() *ListContextDatabaseSyncRunLogsResponseBody
}

type ListContextDatabaseSyncRunLogsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncRunLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncRunLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunLogsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncRunLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncRunLogsResponse) GetBody() *ListContextDatabaseSyncRunLogsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncRunLogsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncRunLogsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncRunLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponse) SetBody(v *ListContextDatabaseSyncRunLogsResponseBody) *ListContextDatabaseSyncRunLogsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncRunLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
