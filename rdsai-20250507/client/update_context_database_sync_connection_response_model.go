// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseSyncConnectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseSyncConnectionResponseBody) *UpdateContextDatabaseSyncConnectionResponse
	GetBody() *UpdateContextDatabaseSyncConnectionResponseBody
}

type UpdateContextDatabaseSyncConnectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseSyncConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseSyncConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseSyncConnectionResponse) GetBody() *UpdateContextDatabaseSyncConnectionResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseSyncConnectionResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponse) SetStatusCode(v int32) *UpdateContextDatabaseSyncConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponse) SetBody(v *UpdateContextDatabaseSyncConnectionResponseBody) *UpdateContextDatabaseSyncConnectionResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
