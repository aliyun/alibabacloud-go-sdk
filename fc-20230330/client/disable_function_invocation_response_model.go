// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFunctionInvocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableFunctionInvocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableFunctionInvocationResponse
	GetStatusCode() *int32
	SetBody(v *DisableFunctionInvocationResponseBody) *DisableFunctionInvocationResponse
	GetBody() *DisableFunctionInvocationResponseBody
}

type DisableFunctionInvocationResponse struct {
	Headers    map[string]*string                     `json:"headers" xml:"headers"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableFunctionInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableFunctionInvocationResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableFunctionInvocationResponse) GoString() string {
	return s.String()
}

func (s *DisableFunctionInvocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableFunctionInvocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableFunctionInvocationResponse) GetBody() *DisableFunctionInvocationResponseBody {
	return s.Body
}

func (s *DisableFunctionInvocationResponse) SetHeaders(v map[string]*string) *DisableFunctionInvocationResponse {
	s.Headers = v
	return s
}

func (s *DisableFunctionInvocationResponse) SetStatusCode(v int32) *DisableFunctionInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableFunctionInvocationResponse) SetBody(v *DisableFunctionInvocationResponseBody) *DisableFunctionInvocationResponse {
	s.Body = v
	return s
}

func (s *DisableFunctionInvocationResponse) Validate() error {
	return dara.Validate(s)
}
