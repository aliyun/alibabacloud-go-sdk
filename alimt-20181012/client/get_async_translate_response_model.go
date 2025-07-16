// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsyncTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsyncTranslateResponse
	GetStatusCode() *int32
	SetBody(v *GetAsyncTranslateResponseBody) *GetAsyncTranslateResponse
	GetBody() *GetAsyncTranslateResponseBody
}

type GetAsyncTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTranslateResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsyncTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsyncTranslateResponse) GetBody() *GetAsyncTranslateResponseBody {
	return s.Body
}

func (s *GetAsyncTranslateResponse) SetHeaders(v map[string]*string) *GetAsyncTranslateResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTranslateResponse) SetStatusCode(v int32) *GetAsyncTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTranslateResponse) SetBody(v *GetAsyncTranslateResponseBody) *GetAsyncTranslateResponse {
	s.Body = v
	return s
}

func (s *GetAsyncTranslateResponse) Validate() error {
	return dara.Validate(s)
}
