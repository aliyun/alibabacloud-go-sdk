// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterUdfFunctionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterUdfFunctionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterUdfFunctionResponse
	GetStatusCode() *int32
	SetBody(v *RegisterUdfFunctionResponseBody) *RegisterUdfFunctionResponse
	GetBody() *RegisterUdfFunctionResponseBody
}

type RegisterUdfFunctionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterUdfFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterUdfFunctionResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterUdfFunctionResponse) GoString() string {
	return s.String()
}

func (s *RegisterUdfFunctionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterUdfFunctionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterUdfFunctionResponse) GetBody() *RegisterUdfFunctionResponseBody {
	return s.Body
}

func (s *RegisterUdfFunctionResponse) SetHeaders(v map[string]*string) *RegisterUdfFunctionResponse {
	s.Headers = v
	return s
}

func (s *RegisterUdfFunctionResponse) SetStatusCode(v int32) *RegisterUdfFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterUdfFunctionResponse) SetBody(v *RegisterUdfFunctionResponseBody) *RegisterUdfFunctionResponse {
	s.Body = v
	return s
}

func (s *RegisterUdfFunctionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
