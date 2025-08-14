// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdInsertionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdInsertionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdInsertionResponse
	GetStatusCode() *int32
	SetBody(v *GetAdInsertionResponseBody) *GetAdInsertionResponse
	GetBody() *GetAdInsertionResponseBody
}

type GetAdInsertionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdInsertionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdInsertionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdInsertionResponse) GoString() string {
	return s.String()
}

func (s *GetAdInsertionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdInsertionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdInsertionResponse) GetBody() *GetAdInsertionResponseBody {
	return s.Body
}

func (s *GetAdInsertionResponse) SetHeaders(v map[string]*string) *GetAdInsertionResponse {
	s.Headers = v
	return s
}

func (s *GetAdInsertionResponse) SetStatusCode(v int32) *GetAdInsertionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdInsertionResponse) SetBody(v *GetAdInsertionResponseBody) *GetAdInsertionResponse {
	s.Body = v
	return s
}

func (s *GetAdInsertionResponse) Validate() error {
	return dara.Validate(s)
}
