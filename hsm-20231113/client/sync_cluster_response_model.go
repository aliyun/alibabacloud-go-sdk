// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncClusterResponse
	GetStatusCode() *int32
	SetBody(v *SyncClusterResponseBody) *SyncClusterResponse
	GetBody() *SyncClusterResponseBody
}

type SyncClusterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncClusterResponse) GoString() string {
	return s.String()
}

func (s *SyncClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncClusterResponse) GetBody() *SyncClusterResponseBody {
	return s.Body
}

func (s *SyncClusterResponse) SetHeaders(v map[string]*string) *SyncClusterResponse {
	s.Headers = v
	return s
}

func (s *SyncClusterResponse) SetStatusCode(v int32) *SyncClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncClusterResponse) SetBody(v *SyncClusterResponseBody) *SyncClusterResponse {
	s.Body = v
	return s
}

func (s *SyncClusterResponse) Validate() error {
	return dara.Validate(s)
}
