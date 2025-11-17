// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserInfoByAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserInfoByAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserInfoByAccountResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserInfoByAccountResponseBody) *QueryUserInfoByAccountResponse
	GetBody() *QueryUserInfoByAccountResponseBody
}

type QueryUserInfoByAccountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserInfoByAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserInfoByAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserInfoByAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryUserInfoByAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserInfoByAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserInfoByAccountResponse) GetBody() *QueryUserInfoByAccountResponseBody {
	return s.Body
}

func (s *QueryUserInfoByAccountResponse) SetHeaders(v map[string]*string) *QueryUserInfoByAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryUserInfoByAccountResponse) SetStatusCode(v int32) *QueryUserInfoByAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserInfoByAccountResponse) SetBody(v *QueryUserInfoByAccountResponseBody) *QueryUserInfoByAccountResponse {
	s.Body = v
	return s
}

func (s *QueryUserInfoByAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
