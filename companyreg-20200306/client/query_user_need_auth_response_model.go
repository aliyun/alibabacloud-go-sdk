// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserNeedAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserNeedAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserNeedAuthResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserNeedAuthResponseBody) *QueryUserNeedAuthResponse
	GetBody() *QueryUserNeedAuthResponseBody
}

type QueryUserNeedAuthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserNeedAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserNeedAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserNeedAuthResponse) GoString() string {
	return s.String()
}

func (s *QueryUserNeedAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserNeedAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserNeedAuthResponse) GetBody() *QueryUserNeedAuthResponseBody {
	return s.Body
}

func (s *QueryUserNeedAuthResponse) SetHeaders(v map[string]*string) *QueryUserNeedAuthResponse {
	s.Headers = v
	return s
}

func (s *QueryUserNeedAuthResponse) SetStatusCode(v int32) *QueryUserNeedAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserNeedAuthResponse) SetBody(v *QueryUserNeedAuthResponseBody) *QueryUserNeedAuthResponse {
	s.Body = v
	return s
}

func (s *QueryUserNeedAuthResponse) Validate() error {
	return dara.Validate(s)
}
