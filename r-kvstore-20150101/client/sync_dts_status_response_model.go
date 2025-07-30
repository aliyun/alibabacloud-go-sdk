// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDtsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDtsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDtsStatusResponse
	GetStatusCode() *int32
	SetBody(v *SyncDtsStatusResponseBody) *SyncDtsStatusResponse
	GetBody() *SyncDtsStatusResponseBody
}

type SyncDtsStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDtsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDtsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDtsStatusResponse) GoString() string {
	return s.String()
}

func (s *SyncDtsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDtsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDtsStatusResponse) GetBody() *SyncDtsStatusResponseBody {
	return s.Body
}

func (s *SyncDtsStatusResponse) SetHeaders(v map[string]*string) *SyncDtsStatusResponse {
	s.Headers = v
	return s
}

func (s *SyncDtsStatusResponse) SetStatusCode(v int32) *SyncDtsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDtsStatusResponse) SetBody(v *SyncDtsStatusResponseBody) *SyncDtsStatusResponse {
	s.Body = v
	return s
}

func (s *SyncDtsStatusResponse) Validate() error {
	return dara.Validate(s)
}
