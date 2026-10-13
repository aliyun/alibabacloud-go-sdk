// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartContextDatabaseSyncRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartContextDatabaseSyncRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartContextDatabaseSyncRunResponse
	GetStatusCode() *int32
	SetBody(v *StartContextDatabaseSyncRunResponseBody) *StartContextDatabaseSyncRunResponse
	GetBody() *StartContextDatabaseSyncRunResponseBody
}

type StartContextDatabaseSyncRunResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartContextDatabaseSyncRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartContextDatabaseSyncRunResponse) String() string {
	return dara.Prettify(s)
}

func (s StartContextDatabaseSyncRunResponse) GoString() string {
	return s.String()
}

func (s *StartContextDatabaseSyncRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartContextDatabaseSyncRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartContextDatabaseSyncRunResponse) GetBody() *StartContextDatabaseSyncRunResponseBody {
	return s.Body
}

func (s *StartContextDatabaseSyncRunResponse) SetHeaders(v map[string]*string) *StartContextDatabaseSyncRunResponse {
	s.Headers = v
	return s
}

func (s *StartContextDatabaseSyncRunResponse) SetStatusCode(v int32) *StartContextDatabaseSyncRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StartContextDatabaseSyncRunResponse) SetBody(v *StartContextDatabaseSyncRunResponseBody) *StartContextDatabaseSyncRunResponse {
	s.Body = v
	return s
}

func (s *StartContextDatabaseSyncRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
