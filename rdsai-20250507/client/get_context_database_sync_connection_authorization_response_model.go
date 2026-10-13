// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncConnectionAuthorizationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncConnectionAuthorizationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncConnectionAuthorizationResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncConnectionAuthorizationResponseBody) *GetContextDatabaseSyncConnectionAuthorizationResponse
	GetBody() *GetContextDatabaseSyncConnectionAuthorizationResponseBody
}

type GetContextDatabaseSyncConnectionAuthorizationResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncConnectionAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncConnectionAuthorizationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncConnectionAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) GetBody() *GetContextDatabaseSyncConnectionAuthorizationResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncConnectionAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) SetStatusCode(v int32) *GetContextDatabaseSyncConnectionAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) SetBody(v *GetContextDatabaseSyncConnectionAuthorizationResponseBody) *GetContextDatabaseSyncConnectionAuthorizationResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncConnectionAuthorizationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
