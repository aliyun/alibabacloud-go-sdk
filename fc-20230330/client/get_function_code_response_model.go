// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionCodeResponse
	GetStatusCode() *int32
	SetBody(v *OutputFuncCode) *GetFunctionCodeResponse
	GetBody() *OutputFuncCode
}

type GetFunctionCodeResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OutputFuncCode    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionCodeResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionCodeResponse) GetBody() *OutputFuncCode {
	return s.Body
}

func (s *GetFunctionCodeResponse) SetHeaders(v map[string]*string) *GetFunctionCodeResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionCodeResponse) SetStatusCode(v int32) *GetFunctionCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionCodeResponse) SetBody(v *OutputFuncCode) *GetFunctionCodeResponse {
	s.Body = v
	return s
}

func (s *GetFunctionCodeResponse) Validate() error {
	return dara.Validate(s)
}
