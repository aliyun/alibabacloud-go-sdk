// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPopupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendPopupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendPopupResponse
	GetStatusCode() *int32
	SetBody(v *SendPopupResponseBody) *SendPopupResponse
	GetBody() *SendPopupResponseBody
}

type SendPopupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendPopupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendPopupResponse) String() string {
	return dara.Prettify(s)
}

func (s SendPopupResponse) GoString() string {
	return s.String()
}

func (s *SendPopupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendPopupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendPopupResponse) GetBody() *SendPopupResponseBody {
	return s.Body
}

func (s *SendPopupResponse) SetHeaders(v map[string]*string) *SendPopupResponse {
	s.Headers = v
	return s
}

func (s *SendPopupResponse) SetStatusCode(v int32) *SendPopupResponse {
	s.StatusCode = &v
	return s
}

func (s *SendPopupResponse) SetBody(v *SendPopupResponseBody) *SendPopupResponse {
	s.Body = v
	return s
}

func (s *SendPopupResponse) Validate() error {
	return dara.Validate(s)
}
