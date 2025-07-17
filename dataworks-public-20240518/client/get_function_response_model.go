// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionResponseBody) *GetFunctionResponse
	GetBody() *GetFunctionResponseBody
}

type GetFunctionResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionResponse) GetBody() *GetFunctionResponseBody {
	return s.Body
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetStatusCode(v int32) *GetFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResponse) SetBody(v *GetFunctionResponseBody) *GetFunctionResponse {
	s.Body = v
	return s
}

func (s *GetFunctionResponse) Validate() error {
	return dara.Validate(s)
}
