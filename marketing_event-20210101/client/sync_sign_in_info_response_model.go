// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncSignInInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncSignInInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncSignInInfoResponse
	GetStatusCode() *int32
	SetBody(v *SyncSignInInfoResponseBody) *SyncSignInInfoResponse
	GetBody() *SyncSignInInfoResponseBody
}

type SyncSignInInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncSignInInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncSignInInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncSignInInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncSignInInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncSignInInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncSignInInfoResponse) GetBody() *SyncSignInInfoResponseBody {
	return s.Body
}

func (s *SyncSignInInfoResponse) SetHeaders(v map[string]*string) *SyncSignInInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncSignInInfoResponse) SetStatusCode(v int32) *SyncSignInInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncSignInInfoResponse) SetBody(v *SyncSignInInfoResponseBody) *SyncSignInInfoResponse {
	s.Body = v
	return s
}

func (s *SyncSignInInfoResponse) Validate() error {
	return dara.Validate(s)
}
