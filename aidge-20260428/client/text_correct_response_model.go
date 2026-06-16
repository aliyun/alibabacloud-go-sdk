// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextCorrectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TextCorrectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TextCorrectResponse
	GetStatusCode() *int32
	SetBody(v *TextCorrectResponseBody) *TextCorrectResponse
	GetBody() *TextCorrectResponseBody
}

type TextCorrectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextCorrectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextCorrectResponse) String() string {
	return dara.Prettify(s)
}

func (s TextCorrectResponse) GoString() string {
	return s.String()
}

func (s *TextCorrectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TextCorrectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TextCorrectResponse) GetBody() *TextCorrectResponseBody {
	return s.Body
}

func (s *TextCorrectResponse) SetHeaders(v map[string]*string) *TextCorrectResponse {
	s.Headers = v
	return s
}

func (s *TextCorrectResponse) SetStatusCode(v int32) *TextCorrectResponse {
	s.StatusCode = &v
	return s
}

func (s *TextCorrectResponse) SetBody(v *TextCorrectResponseBody) *TextCorrectResponse {
	s.Body = v
	return s
}

func (s *TextCorrectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
