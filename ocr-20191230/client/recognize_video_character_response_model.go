// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCharacterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeVideoCharacterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeVideoCharacterResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeVideoCharacterResponseBody) *RecognizeVideoCharacterResponse
	GetBody() *RecognizeVideoCharacterResponseBody
}

type RecognizeVideoCharacterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeVideoCharacterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeVideoCharacterResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeVideoCharacterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeVideoCharacterResponse) GetBody() *RecognizeVideoCharacterResponseBody {
	return s.Body
}

func (s *RecognizeVideoCharacterResponse) SetHeaders(v map[string]*string) *RecognizeVideoCharacterResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVideoCharacterResponse) SetStatusCode(v int32) *RecognizeVideoCharacterResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVideoCharacterResponse) SetBody(v *RecognizeVideoCharacterResponseBody) *RecognizeVideoCharacterResponse {
	s.Body = v
	return s
}

func (s *RecognizeVideoCharacterResponse) Validate() error {
	return dara.Validate(s)
}
