// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncConnectionResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncConnectionResponseBody) *GetContextDatabaseSyncConnectionResponse
	GetBody() *GetContextDatabaseSyncConnectionResponseBody
}

type GetContextDatabaseSyncConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncConnectionResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncConnectionResponse) GetBody() *GetContextDatabaseSyncConnectionResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncConnectionResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncConnectionResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncConnectionResponse) SetStatusCode(v int32) *GetContextDatabaseSyncConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncConnectionResponse) SetBody(v *GetContextDatabaseSyncConnectionResponseBody) *GetContextDatabaseSyncConnectionResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
