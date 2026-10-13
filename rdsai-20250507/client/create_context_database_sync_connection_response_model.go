// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDatabaseSyncConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDatabaseSyncConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDatabaseSyncConnectionResponseBody) *CreateContextDatabaseSyncConnectionResponse
	GetBody() *CreateContextDatabaseSyncConnectionResponseBody
}

type CreateContextDatabaseSyncConnectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDatabaseSyncConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDatabaseSyncConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDatabaseSyncConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDatabaseSyncConnectionResponse) GetBody() *CreateContextDatabaseSyncConnectionResponseBody {
	return s.Body
}

func (s *CreateContextDatabaseSyncConnectionResponse) SetHeaders(v map[string]*string) *CreateContextDatabaseSyncConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDatabaseSyncConnectionResponse) SetStatusCode(v int32) *CreateContextDatabaseSyncConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDatabaseSyncConnectionResponse) SetBody(v *CreateContextDatabaseSyncConnectionResponseBody) *CreateContextDatabaseSyncConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateContextDatabaseSyncConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
