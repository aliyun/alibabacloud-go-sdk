// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustVariableResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustVariableResponseBody) *CreateCustVariableResponse
	GetBody() *CreateCustVariableResponseBody
}

type CreateCustVariableResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateCustVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustVariableResponse) GetBody() *CreateCustVariableResponseBody {
	return s.Body
}

func (s *CreateCustVariableResponse) SetHeaders(v map[string]*string) *CreateCustVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateCustVariableResponse) SetStatusCode(v int32) *CreateCustVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustVariableResponse) SetBody(v *CreateCustVariableResponseBody) *CreateCustVariableResponse {
	s.Body = v
	return s
}

func (s *CreateCustVariableResponse) Validate() error {
	return dara.Validate(s)
}
