// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRoomStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRoomStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRoomStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryRoomStatusResponseBody) *QueryRoomStatusResponse
	GetBody() *QueryRoomStatusResponseBody
}

type QueryRoomStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRoomStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRoomStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRoomStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryRoomStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRoomStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRoomStatusResponse) GetBody() *QueryRoomStatusResponseBody {
	return s.Body
}

func (s *QueryRoomStatusResponse) SetHeaders(v map[string]*string) *QueryRoomStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryRoomStatusResponse) SetStatusCode(v int32) *QueryRoomStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRoomStatusResponse) SetBody(v *QueryRoomStatusResponseBody) *QueryRoomStatusResponse {
	s.Body = v
	return s
}

func (s *QueryRoomStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
