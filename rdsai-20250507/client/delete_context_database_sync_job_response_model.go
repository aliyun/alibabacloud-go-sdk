// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseSyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextDatabaseSyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextDatabaseSyncJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextDatabaseSyncJobResponseBody) *DeleteContextDatabaseSyncJobResponse
	GetBody() *DeleteContextDatabaseSyncJobResponseBody
}

type DeleteContextDatabaseSyncJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextDatabaseSyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextDatabaseSyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseSyncJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseSyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextDatabaseSyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextDatabaseSyncJobResponse) GetBody() *DeleteContextDatabaseSyncJobResponseBody {
	return s.Body
}

func (s *DeleteContextDatabaseSyncJobResponse) SetHeaders(v map[string]*string) *DeleteContextDatabaseSyncJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextDatabaseSyncJobResponse) SetStatusCode(v int32) *DeleteContextDatabaseSyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextDatabaseSyncJobResponse) SetBody(v *DeleteContextDatabaseSyncJobResponseBody) *DeleteContextDatabaseSyncJobResponse {
	s.Body = v
	return s
}

func (s *DeleteContextDatabaseSyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
