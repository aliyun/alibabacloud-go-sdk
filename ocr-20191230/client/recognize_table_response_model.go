// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeTableResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeTableResponseBody) *RecognizeTableResponse
	GetBody() *RecognizeTableResponseBody
}

type RecognizeTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeTableResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeTableResponse) GoString() string {
	return s.String()
}

func (s *RecognizeTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeTableResponse) GetBody() *RecognizeTableResponseBody {
	return s.Body
}

func (s *RecognizeTableResponse) SetHeaders(v map[string]*string) *RecognizeTableResponse {
	s.Headers = v
	return s
}

func (s *RecognizeTableResponse) SetStatusCode(v int32) *RecognizeTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeTableResponse) SetBody(v *RecognizeTableResponseBody) *RecognizeTableResponse {
	s.Body = v
	return s
}

func (s *RecognizeTableResponse) Validate() error {
	return dara.Validate(s)
}
