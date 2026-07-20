// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerNoteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerNoteListResponse
	GetStatusCode() *int32
	SetBody(v *CustomerNoteListResponseBody) *CustomerNoteListResponse
	GetBody() *CustomerNoteListResponseBody
}

type CustomerNoteListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerNoteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerNoteListResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListResponse) GoString() string {
	return s.String()
}

func (s *CustomerNoteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerNoteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerNoteListResponse) GetBody() *CustomerNoteListResponseBody {
	return s.Body
}

func (s *CustomerNoteListResponse) SetHeaders(v map[string]*string) *CustomerNoteListResponse {
	s.Headers = v
	return s
}

func (s *CustomerNoteListResponse) SetStatusCode(v int32) *CustomerNoteListResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerNoteListResponse) SetBody(v *CustomerNoteListResponseBody) *CustomerNoteListResponse {
	s.Body = v
	return s
}

func (s *CustomerNoteListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
