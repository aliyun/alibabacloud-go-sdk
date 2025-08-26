// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizePdfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizePdfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizePdfResponse
	GetStatusCode() *int32
	SetBody(v *RecognizePdfResponseBody) *RecognizePdfResponse
	GetBody() *RecognizePdfResponseBody
}

type RecognizePdfResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizePdfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizePdfResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizePdfResponse) GoString() string {
	return s.String()
}

func (s *RecognizePdfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizePdfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizePdfResponse) GetBody() *RecognizePdfResponseBody {
	return s.Body
}

func (s *RecognizePdfResponse) SetHeaders(v map[string]*string) *RecognizePdfResponse {
	s.Headers = v
	return s
}

func (s *RecognizePdfResponse) SetStatusCode(v int32) *RecognizePdfResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizePdfResponse) SetBody(v *RecognizePdfResponseBody) *RecognizePdfResponse {
	s.Body = v
	return s
}

func (s *RecognizePdfResponse) Validate() error {
	return dara.Validate(s)
}
