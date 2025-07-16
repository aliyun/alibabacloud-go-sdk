// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoByRoomCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryConferenceInfoByRoomCodeResponse
	GetStatusCode() *int32
	SetBody(v *QueryConferenceInfoByRoomCodeResponseBody) *QueryConferenceInfoByRoomCodeResponse
	GetBody() *QueryConferenceInfoByRoomCodeResponseBody
}

type QueryConferenceInfoByRoomCodeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryConferenceInfoByRoomCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryConferenceInfoByRoomCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoByRoomCodeResponse) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoByRoomCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryConferenceInfoByRoomCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryConferenceInfoByRoomCodeResponse) GetBody() *QueryConferenceInfoByRoomCodeResponseBody {
	return s.Body
}

func (s *QueryConferenceInfoByRoomCodeResponse) SetHeaders(v map[string]*string) *QueryConferenceInfoByRoomCodeResponse {
	s.Headers = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponse) SetStatusCode(v int32) *QueryConferenceInfoByRoomCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponse) SetBody(v *QueryConferenceInfoByRoomCodeResponseBody) *QueryConferenceInfoByRoomCodeResponse {
	s.Body = v
	return s
}

func (s *QueryConferenceInfoByRoomCodeResponse) Validate() error {
	return dara.Validate(s)
}
