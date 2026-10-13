// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncPluginsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncPluginsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncPluginsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncPluginsResponseBody) *ListContextDatabaseSyncPluginsResponse
	GetBody() *ListContextDatabaseSyncPluginsResponseBody
}

type ListContextDatabaseSyncPluginsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncPluginsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncPluginsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncPluginsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncPluginsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncPluginsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncPluginsResponse) GetBody() *ListContextDatabaseSyncPluginsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncPluginsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncPluginsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncPluginsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponse) SetBody(v *ListContextDatabaseSyncPluginsResponseBody) *ListContextDatabaseSyncPluginsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncPluginsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
