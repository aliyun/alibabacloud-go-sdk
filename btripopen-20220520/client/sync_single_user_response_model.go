// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSingleUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncSingleUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncSingleUserResponse
	GetStatusCode() *int32
	SetBody(v *SyncSingleUserResponseBody) *SyncSingleUserResponse
	GetBody() *SyncSingleUserResponseBody
}

type SyncSingleUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncSingleUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncSingleUserResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncSingleUserResponse) GoString() string {
	return s.String()
}

func (s *SyncSingleUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncSingleUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncSingleUserResponse) GetBody() *SyncSingleUserResponseBody {
	return s.Body
}

func (s *SyncSingleUserResponse) SetHeaders(v map[string]*string) *SyncSingleUserResponse {
	s.Headers = v
	return s
}

func (s *SyncSingleUserResponse) SetStatusCode(v int32) *SyncSingleUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncSingleUserResponse) SetBody(v *SyncSingleUserResponseBody) *SyncSingleUserResponse {
	s.Body = v
	return s
}

func (s *SyncSingleUserResponse) Validate() error {
	return dara.Validate(s)
}
