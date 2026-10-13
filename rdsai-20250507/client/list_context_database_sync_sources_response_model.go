// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncSourcesResponseBody) *ListContextDatabaseSyncSourcesResponse
	GetBody() *ListContextDatabaseSyncSourcesResponseBody
}

type ListContextDatabaseSyncSourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncSourcesResponse) GetBody() *ListContextDatabaseSyncSourcesResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncSourcesResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponse) SetStatusCode(v int32) *ListContextDatabaseSyncSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponse) SetBody(v *ListContextDatabaseSyncSourcesResponseBody) *ListContextDatabaseSyncSourcesResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
