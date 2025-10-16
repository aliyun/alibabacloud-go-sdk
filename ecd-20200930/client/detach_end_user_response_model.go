// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEndUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachEndUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachEndUserResponse
	GetStatusCode() *int32
	SetBody(v *DetachEndUserResponseBody) *DetachEndUserResponse
	GetBody() *DetachEndUserResponseBody
}

type DetachEndUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachEndUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachEndUserResponse) GoString() string {
	return s.String()
}

func (s *DetachEndUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachEndUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachEndUserResponse) GetBody() *DetachEndUserResponseBody {
	return s.Body
}

func (s *DetachEndUserResponse) SetHeaders(v map[string]*string) *DetachEndUserResponse {
	s.Headers = v
	return s
}

func (s *DetachEndUserResponse) SetStatusCode(v int32) *DetachEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachEndUserResponse) SetBody(v *DetachEndUserResponseBody) *DetachEndUserResponse {
	s.Body = v
	return s
}

func (s *DetachEndUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
