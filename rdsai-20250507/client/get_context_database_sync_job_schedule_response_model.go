// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncJobScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncJobScheduleResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncJobScheduleResponseBody) *GetContextDatabaseSyncJobScheduleResponse
	GetBody() *GetContextDatabaseSyncJobScheduleResponseBody
}

type GetContextDatabaseSyncJobScheduleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncJobScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncJobScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncJobScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncJobScheduleResponse) GetBody() *GetContextDatabaseSyncJobScheduleResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncJobScheduleResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncJobScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponse) SetStatusCode(v int32) *GetContextDatabaseSyncJobScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponse) SetBody(v *GetContextDatabaseSyncJobScheduleResponseBody) *GetContextDatabaseSyncJobScheduleResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncJobScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
