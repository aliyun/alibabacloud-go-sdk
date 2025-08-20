// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVocabResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVocabResponseBody) *UpdateVocabResponse
	GetBody() *UpdateVocabResponseBody
}

type UpdateVocabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabResponse) GoString() string {
	return s.String()
}

func (s *UpdateVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVocabResponse) GetBody() *UpdateVocabResponseBody {
	return s.Body
}

func (s *UpdateVocabResponse) SetHeaders(v map[string]*string) *UpdateVocabResponse {
	s.Headers = v
	return s
}

func (s *UpdateVocabResponse) SetStatusCode(v int32) *UpdateVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVocabResponse) SetBody(v *UpdateVocabResponseBody) *UpdateVocabResponse {
	s.Body = v
	return s
}

func (s *UpdateVocabResponse) Validate() error {
	return dara.Validate(s)
}
