// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFunctionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFunctionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFunctionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetFunctionInstanceResponseBody) *GetFunctionInstanceResponse
	GetBody() *GetFunctionInstanceResponseBody
}

type GetFunctionInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFunctionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFunctionInstanceResponse) GetBody() *GetFunctionInstanceResponseBody {
	return s.Body
}

func (s *GetFunctionInstanceResponse) SetHeaders(v map[string]*string) *GetFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionInstanceResponse) SetStatusCode(v int32) *GetFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionInstanceResponse) SetBody(v *GetFunctionInstanceResponseBody) *GetFunctionInstanceResponse {
	s.Body = v
	return s
}

func (s *GetFunctionInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
