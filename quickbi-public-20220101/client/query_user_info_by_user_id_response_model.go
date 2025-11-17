// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserInfoByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserInfoByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserInfoByUserIdResponseBody) *QueryUserInfoByUserIdResponse
	GetBody() *QueryUserInfoByUserIdResponseBody
}

type QueryUserInfoByUserIdResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserInfoByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserInfoByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserInfoByUserIdResponse) GetBody() *QueryUserInfoByUserIdResponseBody {
	return s.Body
}

func (s *QueryUserInfoByUserIdResponse) SetHeaders(v map[string]*string) *QueryUserInfoByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoByUserIdResponse) SetStatusCode(v int32) *QueryUserInfoByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoByUserIdResponse) SetBody(v *QueryUserInfoByUserIdResponseBody) *QueryUserInfoByUserIdResponse {
	s.Body = v
	return s
}

func (s *QueryUserInfoByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
