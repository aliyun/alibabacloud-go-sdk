// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MuteAllGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MuteAllGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *MuteAllGroupUserResponseBody) *MuteAllGroupUserResponse
	GetBody() *MuteAllGroupUserResponseBody
}

type MuteAllGroupUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MuteAllGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MuteAllGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s MuteAllGroupUserResponse) GoString() string {
	return s.String()
}

func (s *MuteAllGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MuteAllGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MuteAllGroupUserResponse) GetBody() *MuteAllGroupUserResponseBody {
	return s.Body
}

func (s *MuteAllGroupUserResponse) SetHeaders(v map[string]*string) *MuteAllGroupUserResponse {
	s.Headers = v
	return s
}

func (s *MuteAllGroupUserResponse) SetStatusCode(v int32) *MuteAllGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *MuteAllGroupUserResponse) SetBody(v *MuteAllGroupUserResponseBody) *MuteAllGroupUserResponse {
	s.Body = v
	return s
}

func (s *MuteAllGroupUserResponse) Validate() error {
	return dara.Validate(s)
}
