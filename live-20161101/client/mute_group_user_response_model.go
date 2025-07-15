// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MuteGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MuteGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *MuteGroupUserResponseBody) *MuteGroupUserResponse
	GetBody() *MuteGroupUserResponseBody
}

type MuteGroupUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MuteGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s MuteGroupUserResponse) GoString() string {
	return s.String()
}

func (s *MuteGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MuteGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MuteGroupUserResponse) GetBody() *MuteGroupUserResponseBody {
	return s.Body
}

func (s *MuteGroupUserResponse) SetHeaders(v map[string]*string) *MuteGroupUserResponse {
	s.Headers = v
	return s
}

func (s *MuteGroupUserResponse) SetStatusCode(v int32) *MuteGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteGroupUserResponse) SetBody(v *MuteGroupUserResponseBody) *MuteGroupUserResponse {
	s.Body = v
	return s
}

func (s *MuteGroupUserResponse) Validate() error {
	return dara.Validate(s)
}
