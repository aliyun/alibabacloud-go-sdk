// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptVariablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptVariablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptVariablesResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptVariablesResponseBody) *ListScriptVariablesResponse
	GetBody() *ListScriptVariablesResponseBody
}

type ListScriptVariablesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptVariablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVariablesResponse) GoString() string {
	return s.String()
}

func (s *ListScriptVariablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptVariablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptVariablesResponse) GetBody() *ListScriptVariablesResponseBody {
	return s.Body
}

func (s *ListScriptVariablesResponse) SetHeaders(v map[string]*string) *ListScriptVariablesResponse {
	s.Headers = v
	return s
}

func (s *ListScriptVariablesResponse) SetStatusCode(v int32) *ListScriptVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptVariablesResponse) SetBody(v *ListScriptVariablesResponseBody) *ListScriptVariablesResponse {
	s.Body = v
	return s
}

func (s *ListScriptVariablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
