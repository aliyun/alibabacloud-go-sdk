// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseSyncSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDatabaseSyncSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDatabaseSyncSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDatabaseSyncSourceResponseBody) *CreateContextDatabaseSyncSourceResponse
	GetBody() *CreateContextDatabaseSyncSourceResponseBody
}

type CreateContextDatabaseSyncSourceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDatabaseSyncSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDatabaseSyncSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseSyncSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseSyncSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDatabaseSyncSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDatabaseSyncSourceResponse) GetBody() *CreateContextDatabaseSyncSourceResponseBody {
	return s.Body
}

func (s *CreateContextDatabaseSyncSourceResponse) SetHeaders(v map[string]*string) *CreateContextDatabaseSyncSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponse) SetStatusCode(v int32) *CreateContextDatabaseSyncSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponse) SetBody(v *CreateContextDatabaseSyncSourceResponseBody) *CreateContextDatabaseSyncSourceResponse {
	s.Body = v
	return s
}

func (s *CreateContextDatabaseSyncSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
