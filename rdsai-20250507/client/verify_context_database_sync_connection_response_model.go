// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyContextDatabaseSyncConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyContextDatabaseSyncConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyContextDatabaseSyncConnectionResponse
	GetStatusCode() *int32
	SetBody(v *VerifyContextDatabaseSyncConnectionResponseBody) *VerifyContextDatabaseSyncConnectionResponse
	GetBody() *VerifyContextDatabaseSyncConnectionResponseBody
}

type VerifyContextDatabaseSyncConnectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyContextDatabaseSyncConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyContextDatabaseSyncConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyContextDatabaseSyncConnectionResponse) GoString() string {
	return s.String()
}

func (s *VerifyContextDatabaseSyncConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyContextDatabaseSyncConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyContextDatabaseSyncConnectionResponse) GetBody() *VerifyContextDatabaseSyncConnectionResponseBody {
	return s.Body
}

func (s *VerifyContextDatabaseSyncConnectionResponse) SetHeaders(v map[string]*string) *VerifyContextDatabaseSyncConnectionResponse {
	s.Headers = v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponse) SetStatusCode(v int32) *VerifyContextDatabaseSyncConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponse) SetBody(v *VerifyContextDatabaseSyncConnectionResponseBody) *VerifyContextDatabaseSyncConnectionResponse {
	s.Body = v
	return s
}

func (s *VerifyContextDatabaseSyncConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
