// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErrorResponse
	GetStatusCode() *int32
	SetBody(v *GetErrorResponseBody) *GetErrorResponse
	GetBody() *GetErrorResponseBody
}

type GetErrorResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErrorResponse) GoString() string {
	return s.String()
}

func (s *GetErrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErrorResponse) GetBody() *GetErrorResponseBody {
	return s.Body
}

func (s *GetErrorResponse) SetHeaders(v map[string]*string) *GetErrorResponse {
	s.Headers = v
	return s
}

func (s *GetErrorResponse) SetStatusCode(v int32) *GetErrorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorResponse) SetBody(v *GetErrorResponseBody) *GetErrorResponse {
	s.Body = v
	return s
}

func (s *GetErrorResponse) Validate() error {
	return dara.Validate(s)
}
