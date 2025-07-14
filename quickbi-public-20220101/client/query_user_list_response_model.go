// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserListResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserListResponseBody) *QueryUserListResponse
	GetBody() *QueryUserListResponseBody
}

type QueryUserListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserListResponse) GetBody() *QueryUserListResponseBody {
	return s.Body
}

func (s *QueryUserListResponse) SetHeaders(v map[string]*string) *QueryUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserListResponse) SetStatusCode(v int32) *QueryUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserListResponse) SetBody(v *QueryUserListResponseBody) *QueryUserListResponse {
	s.Body = v
	return s
}

func (s *QueryUserListResponse) Validate() error {
	return dara.Validate(s)
}
