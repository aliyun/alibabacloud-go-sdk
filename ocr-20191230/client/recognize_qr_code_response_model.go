// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeQrCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeQrCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeQrCodeResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeQrCodeResponseBody) *RecognizeQrCodeResponse
	GetBody() *RecognizeQrCodeResponseBody
}

type RecognizeQrCodeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeQrCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeQrCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeQrCodeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeQrCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeQrCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeQrCodeResponse) GetBody() *RecognizeQrCodeResponseBody {
	return s.Body
}

func (s *RecognizeQrCodeResponse) SetHeaders(v map[string]*string) *RecognizeQrCodeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeQrCodeResponse) SetStatusCode(v int32) *RecognizeQrCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeQrCodeResponse) SetBody(v *RecognizeQrCodeResponseBody) *RecognizeQrCodeResponse {
	s.Body = v
	return s
}

func (s *RecognizeQrCodeResponse) Validate() error {
	return dara.Validate(s)
}
