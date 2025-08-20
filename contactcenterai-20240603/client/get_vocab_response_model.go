// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVocabResponse
	GetStatusCode() *int32
	SetBody(v *GetVocabResponseBody) *GetVocabResponse
	GetBody() *GetVocabResponseBody
}

type GetVocabResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVocabResponse) GoString() string {
	return s.String()
}

func (s *GetVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVocabResponse) GetBody() *GetVocabResponseBody {
	return s.Body
}

func (s *GetVocabResponse) SetHeaders(v map[string]*string) *GetVocabResponse {
	s.Headers = v
	return s
}

func (s *GetVocabResponse) SetStatusCode(v int32) *GetVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVocabResponse) SetBody(v *GetVocabResponseBody) *GetVocabResponse {
	s.Body = v
	return s
}

func (s *GetVocabResponse) Validate() error {
	return dara.Validate(s)
}
