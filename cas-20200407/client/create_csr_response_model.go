// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCsrResponse
	GetStatusCode() *int32
	SetBody(v *CreateCsrResponseBody) *CreateCsrResponse
	GetBody() *CreateCsrResponseBody
}

type CreateCsrResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCsrResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCsrResponse) GoString() string {
	return s.String()
}

func (s *CreateCsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCsrResponse) GetBody() *CreateCsrResponseBody {
	return s.Body
}

func (s *CreateCsrResponse) SetHeaders(v map[string]*string) *CreateCsrResponse {
	s.Headers = v
	return s
}

func (s *CreateCsrResponse) SetStatusCode(v int32) *CreateCsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCsrResponse) SetBody(v *CreateCsrResponseBody) *CreateCsrResponse {
	s.Body = v
	return s
}

func (s *CreateCsrResponse) Validate() error {
	return dara.Validate(s)
}
