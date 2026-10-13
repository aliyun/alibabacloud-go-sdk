// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncSourceSelectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncSourceSelectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncSourceSelectionResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncSourceSelectionResponseBody) *GetContextDatabaseSyncSourceSelectionResponse
	GetBody() *GetContextDatabaseSyncSourceSelectionResponseBody
}

type GetContextDatabaseSyncSourceSelectionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncSourceSelectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncSourceSelectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceSelectionResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) GetBody() *GetContextDatabaseSyncSourceSelectionResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncSourceSelectionResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) SetStatusCode(v int32) *GetContextDatabaseSyncSourceSelectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) SetBody(v *GetContextDatabaseSyncSourceSelectionResponseBody) *GetContextDatabaseSyncSourceSelectionResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncSourceSelectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
