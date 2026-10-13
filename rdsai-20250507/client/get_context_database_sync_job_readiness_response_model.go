// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncJobReadinessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncJobReadinessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncJobReadinessResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncJobReadinessResponseBody) *GetContextDatabaseSyncJobReadinessResponse
	GetBody() *GetContextDatabaseSyncJobReadinessResponseBody
}

type GetContextDatabaseSyncJobReadinessResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncJobReadinessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncJobReadinessResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncJobReadinessResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncJobReadinessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncJobReadinessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncJobReadinessResponse) GetBody() *GetContextDatabaseSyncJobReadinessResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncJobReadinessResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncJobReadinessResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponse) SetStatusCode(v int32) *GetContextDatabaseSyncJobReadinessResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponse) SetBody(v *GetContextDatabaseSyncJobReadinessResponseBody) *GetContextDatabaseSyncJobReadinessResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncJobReadinessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
