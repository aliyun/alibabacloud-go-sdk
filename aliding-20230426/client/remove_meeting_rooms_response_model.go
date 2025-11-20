// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMeetingRoomsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveMeetingRoomsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveMeetingRoomsResponse
	GetStatusCode() *int32
	SetBody(v *RemoveMeetingRoomsResponseBody) *RemoveMeetingRoomsResponse
	GetBody() *RemoveMeetingRoomsResponseBody
}

type RemoveMeetingRoomsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMeetingRoomsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMeetingRoomsResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveMeetingRoomsResponse) GoString() string {
	return s.String()
}

func (s *RemoveMeetingRoomsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveMeetingRoomsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveMeetingRoomsResponse) GetBody() *RemoveMeetingRoomsResponseBody {
	return s.Body
}

func (s *RemoveMeetingRoomsResponse) SetHeaders(v map[string]*string) *RemoveMeetingRoomsResponse {
	s.Headers = v
	return s
}

func (s *RemoveMeetingRoomsResponse) SetStatusCode(v int32) *RemoveMeetingRoomsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMeetingRoomsResponse) SetBody(v *RemoveMeetingRoomsResponseBody) *RemoveMeetingRoomsResponse {
	s.Body = v
	return s
}

func (s *RemoveMeetingRoomsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
