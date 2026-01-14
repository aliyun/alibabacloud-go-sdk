// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserByMobileAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserByMobileAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserByMobileAccountResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserByMobileAccountResponseBody) *QueryUserByMobileAccountResponse
	GetBody() *QueryUserByMobileAccountResponseBody
}

type QueryUserByMobileAccountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserByMobileAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserByMobileAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserByMobileAccountResponse) GoString() string {
	return s.String()
}

func (s *QueryUserByMobileAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserByMobileAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserByMobileAccountResponse) GetBody() *QueryUserByMobileAccountResponseBody {
	return s.Body
}

func (s *QueryUserByMobileAccountResponse) SetHeaders(v map[string]*string) *QueryUserByMobileAccountResponse {
	s.Headers = v
	return s
}

func (s *QueryUserByMobileAccountResponse) SetStatusCode(v int32) *QueryUserByMobileAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserByMobileAccountResponse) SetBody(v *QueryUserByMobileAccountResponseBody) *QueryUserByMobileAccountResponse {
	s.Body = v
	return s
}

func (s *QueryUserByMobileAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
