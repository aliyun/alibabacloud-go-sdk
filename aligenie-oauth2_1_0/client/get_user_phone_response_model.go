// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPhoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserPhoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserPhoneResponse
	GetStatusCode() *int32
	SetBody(v *GetUserPhoneResponseBody) *GetUserPhoneResponse
	GetBody() *GetUserPhoneResponseBody
}

type GetUserPhoneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserPhoneResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserPhoneResponse) GoString() string {
	return s.String()
}

func (s *GetUserPhoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserPhoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserPhoneResponse) GetBody() *GetUserPhoneResponseBody {
	return s.Body
}

func (s *GetUserPhoneResponse) SetHeaders(v map[string]*string) *GetUserPhoneResponse {
	s.Headers = v
	return s
}

func (s *GetUserPhoneResponse) SetStatusCode(v int32) *GetUserPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserPhoneResponse) SetBody(v *GetUserPhoneResponseBody) *GetUserPhoneResponse {
	s.Body = v
	return s
}

func (s *GetUserPhoneResponse) Validate() error {
	return dara.Validate(s)
}
