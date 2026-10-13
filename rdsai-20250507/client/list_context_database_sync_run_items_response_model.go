// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseSyncRunItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextDatabaseSyncRunItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextDatabaseSyncRunItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListContextDatabaseSyncRunItemsResponseBody) *ListContextDatabaseSyncRunItemsResponse
	GetBody() *ListContextDatabaseSyncRunItemsResponseBody
}

type ListContextDatabaseSyncRunItemsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextDatabaseSyncRunItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextDatabaseSyncRunItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseSyncRunItemsResponse) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseSyncRunItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextDatabaseSyncRunItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextDatabaseSyncRunItemsResponse) GetBody() *ListContextDatabaseSyncRunItemsResponseBody {
	return s.Body
}

func (s *ListContextDatabaseSyncRunItemsResponse) SetHeaders(v map[string]*string) *ListContextDatabaseSyncRunItemsResponse {
	s.Headers = v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponse) SetStatusCode(v int32) *ListContextDatabaseSyncRunItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponse) SetBody(v *ListContextDatabaseSyncRunItemsResponseBody) *ListContextDatabaseSyncRunItemsResponse {
	s.Body = v
	return s
}

func (s *ListContextDatabaseSyncRunItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
