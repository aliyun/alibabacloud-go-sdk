// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeBankCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeBankCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeBankCardResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse
	GetBody() *RecognizeBankCardResponseBody
}

type RecognizeBankCardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeBankCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeBankCardResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeBankCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeBankCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeBankCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeBankCardResponse) GetBody() *RecognizeBankCardResponseBody {
	return s.Body
}

func (s *RecognizeBankCardResponse) SetHeaders(v map[string]*string) *RecognizeBankCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeBankCardResponse) SetStatusCode(v int32) *RecognizeBankCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeBankCardResponse) SetBody(v *RecognizeBankCardResponseBody) *RecognizeBankCardResponse {
	s.Body = v
	return s
}

func (s *RecognizeBankCardResponse) Validate() error {
	return dara.Validate(s)
}
