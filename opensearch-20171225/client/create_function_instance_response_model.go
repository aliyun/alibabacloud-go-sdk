// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFunctionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFunctionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateFunctionInstanceResponseBody) *CreateFunctionInstanceResponse
	GetBody() *CreateFunctionInstanceResponseBody
}

type CreateFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFunctionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFunctionInstanceResponse) GetBody() *CreateFunctionInstanceResponseBody {
	return s.Body
}

func (s *CreateFunctionInstanceResponse) SetHeaders(v map[string]*string) *CreateFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionInstanceResponse) SetStatusCode(v int32) *CreateFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionInstanceResponse) SetBody(v *CreateFunctionInstanceResponseBody) *CreateFunctionInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateFunctionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
