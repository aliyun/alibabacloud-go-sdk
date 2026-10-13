// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncJobScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseSyncJobScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseSyncJobScheduleResponseBody) *UpdateContextDatabaseSyncJobScheduleResponse
	GetBody() *UpdateContextDatabaseSyncJobScheduleResponseBody
}

type UpdateContextDatabaseSyncJobScheduleResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseSyncJobScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseSyncJobScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) GetBody() *UpdateContextDatabaseSyncJobScheduleResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncJobScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) SetStatusCode(v int32) *UpdateContextDatabaseSyncJobScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) SetBody(v *UpdateContextDatabaseSyncJobScheduleResponseBody) *UpdateContextDatabaseSyncJobScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseSyncJobScheduleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
