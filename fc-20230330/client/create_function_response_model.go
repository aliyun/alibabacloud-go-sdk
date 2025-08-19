// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFunctionResponse
	GetStatusCode() *int32
	SetBody(v *Function) *CreateFunctionResponse
	GetBody() *Function
}

type CreateFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Function          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFunctionResponse) GetBody() *Function {
	return s.Body
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetStatusCode(v int32) *CreateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *Function) *CreateFunctionResponse {
	s.Body = v
	return s
}

func (s *CreateFunctionResponse) Validate() error {
	return dara.Validate(s)
}
