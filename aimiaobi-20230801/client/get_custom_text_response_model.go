// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomTextResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomTextResponseBody) *GetCustomTextResponse
	GetBody() *GetCustomTextResponseBody
}

type GetCustomTextResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomTextResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTextResponse) GoString() string {
	return s.String()
}

func (s *GetCustomTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomTextResponse) GetBody() *GetCustomTextResponseBody {
	return s.Body
}

func (s *GetCustomTextResponse) SetHeaders(v map[string]*string) *GetCustomTextResponse {
	s.Headers = v
	return s
}

func (s *GetCustomTextResponse) SetStatusCode(v int32) *GetCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomTextResponse) SetBody(v *GetCustomTextResponseBody) *GetCustomTextResponse {
	s.Body = v
	return s
}

func (s *GetCustomTextResponse) Validate() error {
	return dara.Validate(s)
}
