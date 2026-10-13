// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDatabaseSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDatabaseSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDatabaseSyncJobResponseBody) *CreateContextDatabaseSyncJobResponse
	GetBody() *CreateContextDatabaseSyncJobResponseBody
}

type CreateContextDatabaseSyncJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDatabaseSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDatabaseSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncJobResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDatabaseSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDatabaseSyncJobResponse) GetBody() *CreateContextDatabaseSyncJobResponseBody {
	return s.Body
}

func (s *CreateContextDatabaseSyncJobResponse) SetHeaders(v map[string]*string) *CreateContextDatabaseSyncJobResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDatabaseSyncJobResponse) SetStatusCode(v int32) *CreateContextDatabaseSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDatabaseSyncJobResponse) SetBody(v *CreateContextDatabaseSyncJobResponseBody) *CreateContextDatabaseSyncJobResponse {
	s.Body = v
	return s
}

func (s *CreateContextDatabaseSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
