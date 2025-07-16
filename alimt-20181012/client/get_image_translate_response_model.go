// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetImageTranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetImageTranslateResponse
	GetStatusCode() *int32
	SetBody(v *GetImageTranslateResponseBody) *GetImageTranslateResponse
	GetBody() *GetImageTranslateResponseBody
}

type GetImageTranslateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageTranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageTranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetImageTranslateResponse) GoString() string {
	return s.String()
}

func (s *GetImageTranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetImageTranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetImageTranslateResponse) GetBody() *GetImageTranslateResponseBody {
	return s.Body
}

func (s *GetImageTranslateResponse) SetHeaders(v map[string]*string) *GetImageTranslateResponse {
	s.Headers = v
	return s
}

func (s *GetImageTranslateResponse) SetStatusCode(v int32) *GetImageTranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageTranslateResponse) SetBody(v *GetImageTranslateResponseBody) *GetImageTranslateResponse {
	s.Body = v
	return s
}

func (s *GetImageTranslateResponse) Validate() error {
	return dara.Validate(s)
}
