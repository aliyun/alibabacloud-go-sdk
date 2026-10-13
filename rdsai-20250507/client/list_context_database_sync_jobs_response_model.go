// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncJobsResponseBody) *ListContextDatabaseSyncJobsResponse
	GetBody() *ListContextDatabaseSyncJobsResponseBody
}

type ListContextDatabaseSyncJobsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncJobsResponse) GetBody() *ListContextDatabaseSyncJobsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncJobsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncJobsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncJobsResponse) SetBody(v *ListContextDatabaseSyncJobsResponseBody) *ListContextDatabaseSyncJobsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
