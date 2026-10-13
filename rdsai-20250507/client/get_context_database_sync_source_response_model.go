// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextDatabaseSyncSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextDatabaseSyncSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextDatabaseSyncSourceResponse
	GetStatusCode() *int32
	SetBody(v *GetContextDatabaseSyncSourceResponseBody) *GetContextDatabaseSyncSourceResponse
	GetBody() *GetContextDatabaseSyncSourceResponseBody
}

type GetContextDatabaseSyncSourceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextDatabaseSyncSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextDatabaseSyncSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextDatabaseSyncSourceResponse) GoString() string {
	return s.String()
}

func (s *GetContextDatabaseSyncSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextDatabaseSyncSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextDatabaseSyncSourceResponse) GetBody() *GetContextDatabaseSyncSourceResponseBody {
	return s.Body
}

func (s *GetContextDatabaseSyncSourceResponse) SetHeaders(v map[string]*string) *GetContextDatabaseSyncSourceResponse {
	s.Headers = v
	return s
}

func (s *GetContextDatabaseSyncSourceResponse) SetStatusCode(v int32) *GetContextDatabaseSyncSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextDatabaseSyncSourceResponse) SetBody(v *GetContextDatabaseSyncSourceResponseBody) *GetContextDatabaseSyncSourceResponse {
	s.Body = v
	return s
}

func (s *GetContextDatabaseSyncSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
