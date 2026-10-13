// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseSyncConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextDatabaseSyncConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextDatabaseSyncConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextDatabaseSyncConnectionResponseBody) *DeleteContextDatabaseSyncConnectionResponse
	GetBody() *DeleteContextDatabaseSyncConnectionResponseBody
}

type DeleteContextDatabaseSyncConnectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextDatabaseSyncConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextDatabaseSyncConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseSyncConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseSyncConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextDatabaseSyncConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextDatabaseSyncConnectionResponse) GetBody() *DeleteContextDatabaseSyncConnectionResponseBody {
	return s.Body
}

func (s *DeleteContextDatabaseSyncConnectionResponse) SetHeaders(v map[string]*string) *DeleteContextDatabaseSyncConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextDatabaseSyncConnectionResponse) SetStatusCode(v int32) *DeleteContextDatabaseSyncConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextDatabaseSyncConnectionResponse) SetBody(v *DeleteContextDatabaseSyncConnectionResponseBody) *DeleteContextDatabaseSyncConnectionResponse {
	s.Body = v
	return s
}

func (s *DeleteContextDatabaseSyncConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
