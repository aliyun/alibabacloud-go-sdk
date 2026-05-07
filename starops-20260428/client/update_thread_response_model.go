// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateThreadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateThreadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateThreadResponse
	GetStatusCode() *int32
	SetBody(v *UpdateThreadResponseBody) *UpdateThreadResponse
	GetBody() *UpdateThreadResponseBody
}

type UpdateThreadResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateThreadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateThreadResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateThreadResponse) GoString() string {
	return s.String()
}

func (s *UpdateThreadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateThreadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateThreadResponse) GetBody() *UpdateThreadResponseBody {
	return s.Body
}

func (s *UpdateThreadResponse) SetHeaders(v map[string]*string) *UpdateThreadResponse {
	s.Headers = v
	return s
}

func (s *UpdateThreadResponse) SetStatusCode(v int32) *UpdateThreadResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateThreadResponse) SetBody(v *UpdateThreadResponseBody) *UpdateThreadResponse {
	s.Body = v
	return s
}

func (s *UpdateThreadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
