// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeIdentityCardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeIdentityCardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeIdentityCardResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeIdentityCardResponseBody) *RecognizeIdentityCardResponse
	GetBody() *RecognizeIdentityCardResponseBody
}

type RecognizeIdentityCardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeIdentityCardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIdentityCardResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeIdentityCardResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIdentityCardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeIdentityCardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeIdentityCardResponse) GetBody() *RecognizeIdentityCardResponseBody {
	return s.Body
}

func (s *RecognizeIdentityCardResponse) SetHeaders(v map[string]*string) *RecognizeIdentityCardResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIdentityCardResponse) SetStatusCode(v int32) *RecognizeIdentityCardResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIdentityCardResponse) SetBody(v *RecognizeIdentityCardResponseBody) *RecognizeIdentityCardResponse {
	s.Body = v
	return s
}

func (s *RecognizeIdentityCardResponse) Validate() error {
	return dara.Validate(s)
}
