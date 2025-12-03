// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVariableResponse
	GetStatusCode() *int32
	SetBody(v *GetVariableResponseBody) *GetVariableResponse
	GetBody() *GetVariableResponseBody
}

type GetVariableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVariableResponse) GoString() string {
	return s.String()
}

func (s *GetVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVariableResponse) GetBody() *GetVariableResponseBody {
	return s.Body
}

func (s *GetVariableResponse) SetHeaders(v map[string]*string) *GetVariableResponse {
	s.Headers = v
	return s
}

func (s *GetVariableResponse) SetStatusCode(v int32) *GetVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVariableResponse) SetBody(v *GetVariableResponseBody) *GetVariableResponse {
	s.Body = v
	return s
}

func (s *GetVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
