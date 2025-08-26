// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeCharacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeCharacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeCharacterResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeCharacterResponseBody) *RecognizeCharacterResponse
	GetBody() *RecognizeCharacterResponseBody
}

type RecognizeCharacterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeCharacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeCharacterResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeCharacterResponse) GoString() string {
	return s.String()
}

func (s *RecognizeCharacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeCharacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeCharacterResponse) GetBody() *RecognizeCharacterResponseBody {
	return s.Body
}

func (s *RecognizeCharacterResponse) SetHeaders(v map[string]*string) *RecognizeCharacterResponse {
	s.Headers = v
	return s
}

func (s *RecognizeCharacterResponse) SetStatusCode(v int32) *RecognizeCharacterResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeCharacterResponse) SetBody(v *RecognizeCharacterResponseBody) *RecognizeCharacterResponse {
	s.Body = v
	return s
}

func (s *RecognizeCharacterResponse) Validate() error {
	return dara.Validate(s)
}
