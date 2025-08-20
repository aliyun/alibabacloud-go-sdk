// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVocabResponse
	GetStatusCode() *int32
	SetBody(v *CreateVocabResponseBody) *CreateVocabResponse
	GetBody() *CreateVocabResponseBody
}

type CreateVocabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabResponse) GoString() string {
	return s.String()
}

func (s *CreateVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVocabResponse) GetBody() *CreateVocabResponseBody {
	return s.Body
}

func (s *CreateVocabResponse) SetHeaders(v map[string]*string) *CreateVocabResponse {
	s.Headers = v
	return s
}

func (s *CreateVocabResponse) SetStatusCode(v int32) *CreateVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVocabResponse) SetBody(v *CreateVocabResponseBody) *CreateVocabResponse {
	s.Body = v
	return s
}

func (s *CreateVocabResponse) Validate() error {
	return dara.Validate(s)
}
