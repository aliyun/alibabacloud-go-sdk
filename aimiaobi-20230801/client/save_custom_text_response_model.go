// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCustomTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveCustomTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveCustomTextResponse
	GetStatusCode() *int32
	SetBody(v *SaveCustomTextResponseBody) *SaveCustomTextResponse
	GetBody() *SaveCustomTextResponseBody
}

type SaveCustomTextResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCustomTextResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveCustomTextResponse) GoString() string {
	return s.String()
}

func (s *SaveCustomTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveCustomTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveCustomTextResponse) GetBody() *SaveCustomTextResponseBody {
	return s.Body
}

func (s *SaveCustomTextResponse) SetHeaders(v map[string]*string) *SaveCustomTextResponse {
	s.Headers = v
	return s
}

func (s *SaveCustomTextResponse) SetStatusCode(v int32) *SaveCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCustomTextResponse) SetBody(v *SaveCustomTextResponseBody) *SaveCustomTextResponse {
	s.Body = v
	return s
}

func (s *SaveCustomTextResponse) Validate() error {
	return dara.Validate(s)
}
