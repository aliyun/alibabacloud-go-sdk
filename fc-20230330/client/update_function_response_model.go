// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFunctionResponse
	GetStatusCode() *int32
	SetBody(v *Function) *UpdateFunctionResponse
	GetBody() *Function
}

type UpdateFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Function          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFunctionResponse) GetBody() *Function {
	return s.Body
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetStatusCode(v int32) *UpdateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *Function) *UpdateFunctionResponse {
	s.Body = v
	return s
}

func (s *UpdateFunctionResponse) Validate() error {
	return dara.Validate(s)
}
