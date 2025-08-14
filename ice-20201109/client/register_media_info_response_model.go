// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterMediaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterMediaInfoResponse
	GetStatusCode() *int32
	SetBody(v *RegisterMediaInfoResponseBody) *RegisterMediaInfoResponse
	GetBody() *RegisterMediaInfoResponseBody
}

type RegisterMediaInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterMediaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterMediaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterMediaInfoResponse) GetBody() *RegisterMediaInfoResponseBody {
	return s.Body
}

func (s *RegisterMediaInfoResponse) SetHeaders(v map[string]*string) *RegisterMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *RegisterMediaInfoResponse) SetStatusCode(v int32) *RegisterMediaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterMediaInfoResponse) SetBody(v *RegisterMediaInfoResponseBody) *RegisterMediaInfoResponse {
	s.Body = v
	return s
}

func (s *RegisterMediaInfoResponse) Validate() error {
	return dara.Validate(s)
}
