// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryFailedContextDatabaseSyncRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryFailedContextDatabaseSyncRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryFailedContextDatabaseSyncRunResponse
	GetStatusCode() *int32
	SetBody(v *RetryFailedContextDatabaseSyncRunResponseBody) *RetryFailedContextDatabaseSyncRunResponse
	GetBody() *RetryFailedContextDatabaseSyncRunResponseBody
}

type RetryFailedContextDatabaseSyncRunResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryFailedContextDatabaseSyncRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryFailedContextDatabaseSyncRunResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryFailedContextDatabaseSyncRunResponse) GoString() string {
	return s.String()
}

func (s *RetryFailedContextDatabaseSyncRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryFailedContextDatabaseSyncRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryFailedContextDatabaseSyncRunResponse) GetBody() *RetryFailedContextDatabaseSyncRunResponseBody {
	return s.Body
}

func (s *RetryFailedContextDatabaseSyncRunResponse) SetHeaders(v map[string]*string) *RetryFailedContextDatabaseSyncRunResponse {
	s.Headers = v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponse) SetStatusCode(v int32) *RetryFailedContextDatabaseSyncRunResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponse) SetBody(v *RetryFailedContextDatabaseSyncRunResponseBody) *RetryFailedContextDatabaseSyncRunResponse {
	s.Body = v
	return s
}

func (s *RetryFailedContextDatabaseSyncRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
