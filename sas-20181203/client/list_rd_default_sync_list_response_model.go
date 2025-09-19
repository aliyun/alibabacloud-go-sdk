// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRdDefaultSyncListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRdDefaultSyncListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRdDefaultSyncListResponse
	GetStatusCode() *int32
	SetBody(v *ListRdDefaultSyncListResponseBody) *ListRdDefaultSyncListResponse
	GetBody() *ListRdDefaultSyncListResponseBody
}

type ListRdDefaultSyncListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRdDefaultSyncListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRdDefaultSyncListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRdDefaultSyncListResponse) GoString() string {
	return s.String()
}

func (s *ListRdDefaultSyncListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRdDefaultSyncListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRdDefaultSyncListResponse) GetBody() *ListRdDefaultSyncListResponseBody {
	return s.Body
}

func (s *ListRdDefaultSyncListResponse) SetHeaders(v map[string]*string) *ListRdDefaultSyncListResponse {
	s.Headers = v
	return s
}

func (s *ListRdDefaultSyncListResponse) SetStatusCode(v int32) *ListRdDefaultSyncListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRdDefaultSyncListResponse) SetBody(v *ListRdDefaultSyncListResponseBody) *ListRdDefaultSyncListResponse {
	s.Body = v
	return s
}

func (s *ListRdDefaultSyncListResponse) Validate() error {
	return dara.Validate(s)
}
