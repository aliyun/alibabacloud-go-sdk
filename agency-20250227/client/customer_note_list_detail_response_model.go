// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteListDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerNoteListDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerNoteListDetailResponse
	GetStatusCode() *int32
	SetBody(v *CustomerNoteListDetailResponseBody) *CustomerNoteListDetailResponse
	GetBody() *CustomerNoteListDetailResponseBody
}

type CustomerNoteListDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerNoteListDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerNoteListDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListDetailResponse) GoString() string {
	return s.String()
}

func (s *CustomerNoteListDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerNoteListDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerNoteListDetailResponse) GetBody() *CustomerNoteListDetailResponseBody {
	return s.Body
}

func (s *CustomerNoteListDetailResponse) SetHeaders(v map[string]*string) *CustomerNoteListDetailResponse {
	s.Headers = v
	return s
}

func (s *CustomerNoteListDetailResponse) SetStatusCode(v int32) *CustomerNoteListDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerNoteListDetailResponse) SetBody(v *CustomerNoteListDetailResponseBody) *CustomerNoteListDetailResponse {
	s.Body = v
	return s
}

func (s *CustomerNoteListDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
