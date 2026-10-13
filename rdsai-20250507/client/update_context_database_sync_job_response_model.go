// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseSyncJobResponseBody) *UpdateContextDatabaseSyncJobResponse
	GetBody() *UpdateContextDatabaseSyncJobResponseBody
}

type UpdateContextDatabaseSyncJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseSyncJobResponse) GetBody() *UpdateContextDatabaseSyncJobResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseSyncJobResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponse) SetStatusCode(v int32) *UpdateContextDatabaseSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponse) SetBody(v *UpdateContextDatabaseSyncJobResponseBody) *UpdateContextDatabaseSyncJobResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
