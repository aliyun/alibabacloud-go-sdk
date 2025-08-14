// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQueryVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQueryVariableResponse
	GetStatusCode() *int32
	SetBody(v *CreateQueryVariableResponseBody) *CreateQueryVariableResponse
	GetBody() *CreateQueryVariableResponseBody
}

type CreateQueryVariableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueryVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueryVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateQueryVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQueryVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQueryVariableResponse) GetBody() *CreateQueryVariableResponseBody {
	return s.Body
}

func (s *CreateQueryVariableResponse) SetHeaders(v map[string]*string) *CreateQueryVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateQueryVariableResponse) SetStatusCode(v int32) *CreateQueryVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueryVariableResponse) SetBody(v *CreateQueryVariableResponseBody) *CreateQueryVariableResponse {
	s.Body = v
	return s
}

func (s *CreateQueryVariableResponse) Validate() error {
	return dara.Validate(s)
}
