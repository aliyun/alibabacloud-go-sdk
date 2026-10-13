// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseSyncSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextDatabaseSyncSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextDatabaseSyncSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextDatabaseSyncSourceResponseBody) *DeleteContextDatabaseSyncSourceResponse
	GetBody() *DeleteContextDatabaseSyncSourceResponseBody
}

type DeleteContextDatabaseSyncSourceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextDatabaseSyncSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextDatabaseSyncSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseSyncSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseSyncSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextDatabaseSyncSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextDatabaseSyncSourceResponse) GetBody() *DeleteContextDatabaseSyncSourceResponseBody {
	return s.Body
}

func (s *DeleteContextDatabaseSyncSourceResponse) SetHeaders(v map[string]*string) *DeleteContextDatabaseSyncSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextDatabaseSyncSourceResponse) SetStatusCode(v int32) *DeleteContextDatabaseSyncSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextDatabaseSyncSourceResponse) SetBody(v *DeleteContextDatabaseSyncSourceResponseBody) *DeleteContextDatabaseSyncSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteContextDatabaseSyncSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
