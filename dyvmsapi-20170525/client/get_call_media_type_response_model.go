// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCallMediaTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCallMediaTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCallMediaTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetCallMediaTypeResponseBody) *GetCallMediaTypeResponse
	GetBody() *GetCallMediaTypeResponseBody
}

type GetCallMediaTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallMediaTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallMediaTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCallMediaTypeResponse) GoString() string {
	return s.String()
}

func (s *GetCallMediaTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCallMediaTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCallMediaTypeResponse) GetBody() *GetCallMediaTypeResponseBody {
	return s.Body
}

func (s *GetCallMediaTypeResponse) SetHeaders(v map[string]*string) *GetCallMediaTypeResponse {
	s.Headers = v
	return s
}

func (s *GetCallMediaTypeResponse) SetStatusCode(v int32) *GetCallMediaTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallMediaTypeResponse) SetBody(v *GetCallMediaTypeResponseBody) *GetCallMediaTypeResponse {
	s.Body = v
	return s
}

func (s *GetCallMediaTypeResponse) Validate() error {
	return dara.Validate(s)
}
