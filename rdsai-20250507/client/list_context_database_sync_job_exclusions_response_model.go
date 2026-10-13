// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobExclusionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncJobExclusionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncJobExclusionsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncJobExclusionsResponseBody) *ListContextDatabaseSyncJobExclusionsResponse
	GetBody() *ListContextDatabaseSyncJobExclusionsResponseBody
}

type ListContextDatabaseSyncJobExclusionsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncJobExclusionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncJobExclusionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobExclusionsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) GetBody() *ListContextDatabaseSyncJobExclusionsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncJobExclusionsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncJobExclusionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) SetBody(v *ListContextDatabaseSyncJobExclusionsResponseBody) *ListContextDatabaseSyncJobExclusionsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncJobExclusionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
