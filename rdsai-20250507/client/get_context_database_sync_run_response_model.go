// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncRunResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncRunResponseBody) *GetContextDatabaseSyncRunResponse
	GetBody() *GetContextDatabaseSyncRunResponseBody
}

type GetContextDatabaseSyncRunResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncRunResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncRunResponse) GetBody() *GetContextDatabaseSyncRunResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncRunResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncRunResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncRunResponse) SetStatusCode(v int32) *GetContextDatabaseSyncRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncRunResponse) SetBody(v *GetContextDatabaseSyncRunResponseBody) *GetContextDatabaseSyncRunResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
