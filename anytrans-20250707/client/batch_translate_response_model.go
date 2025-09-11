// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchTranslateResponse
	GetStatusCode() *int32
	SetBody(v *BatchTranslateResponseBody) *BatchTranslateResponse
	GetBody() *BatchTranslateResponseBody
}

type BatchTranslateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchTranslateResponse) GoString() string {
	return s.String()
}

func (s *BatchTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchTranslateResponse) GetBody() *BatchTranslateResponseBody {
	return s.Body
}

func (s *BatchTranslateResponse) SetHeaders(v map[string]*string) *BatchTranslateResponse {
	s.Headers = v
	return s
}

func (s *BatchTranslateResponse) SetStatusCode(v int32) *BatchTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchTranslateResponse) SetBody(v *BatchTranslateResponseBody) *BatchTranslateResponse {
	s.Body = v
	return s
}

func (s *BatchTranslateResponse) Validate() error {
	return dara.Validate(s)
}
