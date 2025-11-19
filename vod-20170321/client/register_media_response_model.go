// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterMediaResponse
	GetStatusCode() *int32
	SetBody(v *RegisterMediaResponseBody) *RegisterMediaResponse
	GetBody() *RegisterMediaResponseBody
}

type RegisterMediaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaResponse) GoString() string {
	return s.String()
}

func (s *RegisterMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterMediaResponse) GetBody() *RegisterMediaResponseBody {
	return s.Body
}

func (s *RegisterMediaResponse) SetHeaders(v map[string]*string) *RegisterMediaResponse {
	s.Headers = v
	return s
}

func (s *RegisterMediaResponse) SetStatusCode(v int32) *RegisterMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterMediaResponse) SetBody(v *RegisterMediaResponseBody) *RegisterMediaResponse {
	s.Body = v
	return s
}

func (s *RegisterMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
