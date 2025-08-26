// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeLicensePlateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecognizeLicensePlateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecognizeLicensePlateResponse
	GetStatusCode() *int32
	SetBody(v *RecognizeLicensePlateResponseBody) *RecognizeLicensePlateResponse
	GetBody() *RecognizeLicensePlateResponseBody
}

type RecognizeLicensePlateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeLicensePlateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeLicensePlateResponse) String() string {
	return dara.Prettify(s)
}

func (s RecognizeLicensePlateResponse) GoString() string {
	return s.String()
}

func (s *RecognizeLicensePlateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecognizeLicensePlateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecognizeLicensePlateResponse) GetBody() *RecognizeLicensePlateResponseBody {
	return s.Body
}

func (s *RecognizeLicensePlateResponse) SetHeaders(v map[string]*string) *RecognizeLicensePlateResponse {
	s.Headers = v
	return s
}

func (s *RecognizeLicensePlateResponse) SetStatusCode(v int32) *RecognizeLicensePlateResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeLicensePlateResponse) SetBody(v *RecognizeLicensePlateResponseBody) *RecognizeLicensePlateResponse {
	s.Body = v
	return s
}

func (s *RecognizeLicensePlateResponse) Validate() error {
	return dara.Validate(s)
}
