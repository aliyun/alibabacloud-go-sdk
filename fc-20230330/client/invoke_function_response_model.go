// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iInvokeFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InvokeFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InvokeFunctionResponse
	GetStatusCode() *int32
	SetBody(v io.Reader) *InvokeFunctionResponse
	GetBody() io.Reader
}

type InvokeFunctionResponse struct {
	Headers    map[string]*string `json:"headers" xml:"headers"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s InvokeFunctionResponse) GoString() string {
	return s.String()
}

func (s *InvokeFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeFunctionResponse) GetBody() io.Reader {
	return s.Body
}

func (s *InvokeFunctionResponse) SetHeaders(v map[string]*string) *InvokeFunctionResponse {
	s.Headers = v
	return s
}

func (s *InvokeFunctionResponse) SetStatusCode(v int32) *InvokeFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeFunctionResponse) SetBody(v io.Reader) *InvokeFunctionResponse {
	s.Body = v
	return s
}

func (s *InvokeFunctionResponse) Validate() error {
	return dara.Validate(s)
}
