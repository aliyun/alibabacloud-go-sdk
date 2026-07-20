// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteEditResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerNoteEditResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerNoteEditResponse
	GetStatusCode() *int32
	SetBody(v *CustomerNoteEditResponseBody) *CustomerNoteEditResponse
	GetBody() *CustomerNoteEditResponseBody
}

type CustomerNoteEditResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerNoteEditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerNoteEditResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteEditResponse) GoString() string {
	return s.String()
}

func (s *CustomerNoteEditResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerNoteEditResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerNoteEditResponse) GetBody() *CustomerNoteEditResponseBody {
	return s.Body
}

func (s *CustomerNoteEditResponse) SetHeaders(v map[string]*string) *CustomerNoteEditResponse {
	s.Headers = v
	return s
}

func (s *CustomerNoteEditResponse) SetStatusCode(v int32) *CustomerNoteEditResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerNoteEditResponse) SetBody(v *CustomerNoteEditResponseBody) *CustomerNoteEditResponse {
	s.Body = v
	return s
}

func (s *CustomerNoteEditResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
