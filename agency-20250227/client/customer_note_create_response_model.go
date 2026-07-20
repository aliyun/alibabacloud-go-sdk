// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerNoteCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerNoteCreateResponse
	GetStatusCode() *int32
	SetBody(v *CustomerNoteCreateResponseBody) *CustomerNoteCreateResponse
	GetBody() *CustomerNoteCreateResponseBody
}

type CustomerNoteCreateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerNoteCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerNoteCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteCreateResponse) GoString() string {
	return s.String()
}

func (s *CustomerNoteCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerNoteCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerNoteCreateResponse) GetBody() *CustomerNoteCreateResponseBody {
	return s.Body
}

func (s *CustomerNoteCreateResponse) SetHeaders(v map[string]*string) *CustomerNoteCreateResponse {
	s.Headers = v
	return s
}

func (s *CustomerNoteCreateResponse) SetStatusCode(v int32) *CustomerNoteCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerNoteCreateResponse) SetBody(v *CustomerNoteCreateResponseBody) *CustomerNoteCreateResponse {
	s.Body = v
	return s
}

func (s *CustomerNoteCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
