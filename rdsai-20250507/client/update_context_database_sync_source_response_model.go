// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseSyncSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseSyncSourceResponseBody) *UpdateContextDatabaseSyncSourceResponse
	GetBody() *UpdateContextDatabaseSyncSourceResponseBody
}

type UpdateContextDatabaseSyncSourceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseSyncSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseSyncSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseSyncSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseSyncSourceResponse) GetBody() *UpdateContextDatabaseSyncSourceResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseSyncSourceResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseSyncSourceResponse) SetStatusCode(v int32) *UpdateContextDatabaseSyncSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceResponse) SetBody(v *UpdateContextDatabaseSyncSourceResponseBody) *UpdateContextDatabaseSyncSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseSyncSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
