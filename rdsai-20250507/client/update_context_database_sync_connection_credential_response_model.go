// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncConnectionCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncConnectionCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseSyncConnectionCredentialResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseSyncConnectionCredentialResponseBody) *UpdateContextDatabaseSyncConnectionCredentialResponse
	GetBody() *UpdateContextDatabaseSyncConnectionCredentialResponseBody
}

type UpdateContextDatabaseSyncConnectionCredentialResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseSyncConnectionCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseSyncConnectionCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncConnectionCredentialResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) GetBody() *UpdateContextDatabaseSyncConnectionCredentialResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncConnectionCredentialResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) SetStatusCode(v int32) *UpdateContextDatabaseSyncConnectionCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) SetBody(v *UpdateContextDatabaseSyncConnectionCredentialResponseBody) *UpdateContextDatabaseSyncConnectionCredentialResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseSyncConnectionCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
