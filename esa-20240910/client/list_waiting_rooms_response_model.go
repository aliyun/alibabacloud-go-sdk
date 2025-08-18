// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaitingRoomsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWaitingRoomsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWaitingRoomsResponse
	GetStatusCode() *int32
	SetBody(v *ListWaitingRoomsResponseBody) *ListWaitingRoomsResponse
	GetBody() *ListWaitingRoomsResponseBody
}

type ListWaitingRoomsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWaitingRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWaitingRoomsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWaitingRoomsResponse) GoString() string {
	return s.String()
}

func (s *ListWaitingRoomsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWaitingRoomsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWaitingRoomsResponse) GetBody() *ListWaitingRoomsResponseBody {
	return s.Body
}

func (s *ListWaitingRoomsResponse) SetHeaders(v map[string]*string) *ListWaitingRoomsResponse {
	s.Headers = v
	return s
}

func (s *ListWaitingRoomsResponse) SetStatusCode(v int32) *ListWaitingRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWaitingRoomsResponse) SetBody(v *ListWaitingRoomsResponseBody) *ListWaitingRoomsResponse {
	s.Body = v
	return s
}

func (s *ListWaitingRoomsResponse) Validate() error {
	return dara.Validate(s)
}
