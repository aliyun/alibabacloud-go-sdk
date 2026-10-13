// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncSourceTreeNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncSourceTreeNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncSourceTreeNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncSourceTreeNodesResponseBody) *ListContextDatabaseSyncSourceTreeNodesResponse
	GetBody() *ListContextDatabaseSyncSourceTreeNodesResponseBody
}

type ListContextDatabaseSyncSourceTreeNodesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncSourceTreeNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncSourceTreeNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncSourceTreeNodesResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) GetBody() *ListContextDatabaseSyncSourceTreeNodesResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncSourceTreeNodesResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) SetStatusCode(v int32) *ListContextDatabaseSyncSourceTreeNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) SetBody(v *ListContextDatabaseSyncSourceTreeNodesResponseBody) *ListContextDatabaseSyncSourceTreeNodesResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncSourceTreeNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
