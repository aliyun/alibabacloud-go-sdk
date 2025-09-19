// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinWebLockProcessWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinWebLockProcessWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinWebLockProcessWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *JoinWebLockProcessWhiteListResponseBody) *JoinWebLockProcessWhiteListResponse
	GetBody() *JoinWebLockProcessWhiteListResponseBody
}

type JoinWebLockProcessWhiteListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinWebLockProcessWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinWebLockProcessWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinWebLockProcessWhiteListResponse) GoString() string {
	return s.String()
}

func (s *JoinWebLockProcessWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinWebLockProcessWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinWebLockProcessWhiteListResponse) GetBody() *JoinWebLockProcessWhiteListResponseBody {
	return s.Body
}

func (s *JoinWebLockProcessWhiteListResponse) SetHeaders(v map[string]*string) *JoinWebLockProcessWhiteListResponse {
	s.Headers = v
	return s
}

func (s *JoinWebLockProcessWhiteListResponse) SetStatusCode(v int32) *JoinWebLockProcessWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinWebLockProcessWhiteListResponse) SetBody(v *JoinWebLockProcessWhiteListResponseBody) *JoinWebLockProcessWhiteListResponse {
	s.Body = v
	return s
}

func (s *JoinWebLockProcessWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
