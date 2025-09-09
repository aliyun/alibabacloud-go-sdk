// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRTCLiveRoomsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRTCLiveRoomsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRTCLiveRoomsResponse
	GetStatusCode() *int32
	SetBody(v *ListRTCLiveRoomsResponseBody) *ListRTCLiveRoomsResponse
	GetBody() *ListRTCLiveRoomsResponseBody
}

type ListRTCLiveRoomsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRTCLiveRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRTCLiveRoomsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRTCLiveRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListRTCLiveRoomsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRTCLiveRoomsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRTCLiveRoomsResponse) GetBody() *ListRTCLiveRoomsResponseBody {
	return s.Body
}

func (s *ListRTCLiveRoomsResponse) SetHeaders(v map[string]*string) *ListRTCLiveRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListRTCLiveRoomsResponse) SetStatusCode(v int32) *ListRTCLiveRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRTCLiveRoomsResponse) SetBody(v *ListRTCLiveRoomsResponseBody) *ListRTCLiveRoomsResponse {
	s.Body = v
	return s
}

func (s *ListRTCLiveRoomsResponse) Validate() error {
	return dara.Validate(s)
}
