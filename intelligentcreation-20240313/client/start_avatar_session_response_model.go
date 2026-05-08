// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAvatarSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAvatarSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAvatarSessionResponse
	GetStatusCode() *int32
	SetBody(v *StartAvatarSessionResponseBody) *StartAvatarSessionResponse
	GetBody() *StartAvatarSessionResponseBody
}

type StartAvatarSessionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAvatarSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAvatarSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAvatarSessionResponse) GoString() string {
	return s.String()
}

func (s *StartAvatarSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAvatarSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAvatarSessionResponse) GetBody() *StartAvatarSessionResponseBody {
	return s.Body
}

func (s *StartAvatarSessionResponse) SetHeaders(v map[string]*string) *StartAvatarSessionResponse {
	s.Headers = v
	return s
}

func (s *StartAvatarSessionResponse) SetStatusCode(v int32) *StartAvatarSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAvatarSessionResponse) SetBody(v *StartAvatarSessionResponseBody) *StartAvatarSessionResponse {
	s.Body = v
	return s
}

func (s *StartAvatarSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
