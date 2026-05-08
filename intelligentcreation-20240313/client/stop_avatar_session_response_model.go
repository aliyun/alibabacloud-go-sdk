// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAvatarSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAvatarSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAvatarSessionResponse
	GetStatusCode() *int32
	SetBody(v *StopAvatarSessionResponseBody) *StopAvatarSessionResponse
	GetBody() *StopAvatarSessionResponseBody
}

type StopAvatarSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAvatarSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAvatarSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAvatarSessionResponse) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAvatarSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAvatarSessionResponse) GetBody() *StopAvatarSessionResponseBody {
	return s.Body
}

func (s *StopAvatarSessionResponse) SetHeaders(v map[string]*string) *StopAvatarSessionResponse {
	s.Headers = v
	return s
}

func (s *StopAvatarSessionResponse) SetStatusCode(v int32) *StopAvatarSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAvatarSessionResponse) SetBody(v *StopAvatarSessionResponseBody) *StopAvatarSessionResponse {
	s.Body = v
	return s
}

func (s *StopAvatarSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
