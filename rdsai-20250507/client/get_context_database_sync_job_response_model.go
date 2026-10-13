// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncJobResponseBody) *GetContextDatabaseSyncJobResponse
	GetBody() *GetContextDatabaseSyncJobResponseBody
}

type GetContextDatabaseSyncJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncJobResponse) GetBody() *GetContextDatabaseSyncJobResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncJobResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncJobResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncJobResponse) SetStatusCode(v int32) *GetContextDatabaseSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncJobResponse) SetBody(v *GetContextDatabaseSyncJobResponseBody) *GetContextDatabaseSyncJobResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
