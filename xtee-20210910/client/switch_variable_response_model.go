// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchVariableResponse
	GetStatusCode() *int32
	SetBody(v *SwitchVariableResponseBody) *SwitchVariableResponse
	GetBody() *SwitchVariableResponseBody
}

type SwitchVariableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchVariableResponse) GoString() string {
	return s.String()
}

func (s *SwitchVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchVariableResponse) GetBody() *SwitchVariableResponseBody {
	return s.Body
}

func (s *SwitchVariableResponse) SetHeaders(v map[string]*string) *SwitchVariableResponse {
	s.Headers = v
	return s
}

func (s *SwitchVariableResponse) SetStatusCode(v int32) *SwitchVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchVariableResponse) SetBody(v *SwitchVariableResponseBody) *SwitchVariableResponse {
	s.Body = v
	return s
}

func (s *SwitchVariableResponse) Validate() error {
	return dara.Validate(s)
}
