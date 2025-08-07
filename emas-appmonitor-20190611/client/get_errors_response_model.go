// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErrorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErrorsResponse
	GetStatusCode() *int32
	SetBody(v *GetErrorsResponseBody) *GetErrorsResponse
	GetBody() *GetErrorsResponseBody
}

type GetErrorsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErrorsResponse) GoString() string {
	return s.String()
}

func (s *GetErrorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErrorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErrorsResponse) GetBody() *GetErrorsResponseBody {
	return s.Body
}

func (s *GetErrorsResponse) SetHeaders(v map[string]*string) *GetErrorsResponse {
	s.Headers = v
	return s
}

func (s *GetErrorsResponse) SetStatusCode(v int32) *GetErrorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorsResponse) SetBody(v *GetErrorsResponseBody) *GetErrorsResponse {
	s.Body = v
	return s
}

func (s *GetErrorsResponse) Validate() error {
	return dara.Validate(s)
}
