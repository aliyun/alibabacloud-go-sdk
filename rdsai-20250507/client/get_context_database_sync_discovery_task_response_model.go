// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncDiscoveryTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncDiscoveryTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncDiscoveryTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncDiscoveryTaskResponseBody) *GetContextDatabaseSyncDiscoveryTaskResponse
	GetBody() *GetContextDatabaseSyncDiscoveryTaskResponseBody
}

type GetContextDatabaseSyncDiscoveryTaskResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncDiscoveryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncDiscoveryTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncDiscoveryTaskResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) GetBody() *GetContextDatabaseSyncDiscoveryTaskResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncDiscoveryTaskResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) SetStatusCode(v int32) *GetContextDatabaseSyncDiscoveryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) SetBody(v *GetContextDatabaseSyncDiscoveryTaskResponseBody) *GetContextDatabaseSyncDiscoveryTaskResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncDiscoveryTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
