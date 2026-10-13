// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncJobDocumentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncJobDocumentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncJobDocumentsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncJobDocumentsResponseBody) *ListContextDatabaseSyncJobDocumentsResponse
	GetBody() *ListContextDatabaseSyncJobDocumentsResponseBody
}

type ListContextDatabaseSyncJobDocumentsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncJobDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncJobDocumentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncJobDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) GetBody() *ListContextDatabaseSyncJobDocumentsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncJobDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncJobDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) SetBody(v *ListContextDatabaseSyncJobDocumentsResponseBody) *ListContextDatabaseSyncJobDocumentsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncJobDocumentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
