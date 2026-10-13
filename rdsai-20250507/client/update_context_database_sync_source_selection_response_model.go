// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseSyncSourceSelectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncSourceSelectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseSyncSourceSelectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseSyncSourceSelectionResponseBody) *UpdateContextDatabaseSyncSourceSelectionResponse
	GetBody() *UpdateContextDatabaseSyncSourceSelectionResponseBody
}

type UpdateContextDatabaseSyncSourceSelectionResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseSyncSourceSelectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseSyncSourceSelectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseSyncSourceSelectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) GetBody() *UpdateContextDatabaseSyncSourceSelectionResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseSyncSourceSelectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) SetStatusCode(v int32) *UpdateContextDatabaseSyncSourceSelectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) SetBody(v *UpdateContextDatabaseSyncSourceSelectionResponseBody) *UpdateContextDatabaseSyncSourceSelectionResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseSyncSourceSelectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
