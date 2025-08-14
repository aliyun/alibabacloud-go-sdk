// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchQueryVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchQueryVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchQueryVariableResponse
	GetStatusCode() *int32
	SetBody(v *SwitchQueryVariableResponseBody) *SwitchQueryVariableResponse
	GetBody() *SwitchQueryVariableResponseBody
}

type SwitchQueryVariableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchQueryVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchQueryVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchQueryVariableResponse) GoString() string {
	return s.String()
}

func (s *SwitchQueryVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchQueryVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchQueryVariableResponse) GetBody() *SwitchQueryVariableResponseBody {
	return s.Body
}

func (s *SwitchQueryVariableResponse) SetHeaders(v map[string]*string) *SwitchQueryVariableResponse {
	s.Headers = v
	return s
}

func (s *SwitchQueryVariableResponse) SetStatusCode(v int32) *SwitchQueryVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchQueryVariableResponse) SetBody(v *SwitchQueryVariableResponseBody) *SwitchQueryVariableResponse {
	s.Body = v
	return s
}

func (s *SwitchQueryVariableResponse) Validate() error {
	return dara.Validate(s)
}
