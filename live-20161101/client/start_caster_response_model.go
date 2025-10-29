// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCasterResponse
	GetStatusCode() *int32
	SetBody(v *StartCasterResponseBody) *StartCasterResponse
	GetBody() *StartCasterResponseBody
}

type StartCasterResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCasterResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCasterResponse) GoString() string {
	return s.String()
}

func (s *StartCasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCasterResponse) GetBody() *StartCasterResponseBody {
	return s.Body
}

func (s *StartCasterResponse) SetHeaders(v map[string]*string) *StartCasterResponse {
	s.Headers = v
	return s
}

func (s *StartCasterResponse) SetStatusCode(v int32) *StartCasterResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCasterResponse) SetBody(v *StartCasterResponseBody) *StartCasterResponse {
	s.Body = v
	return s
}

func (s *StartCasterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
