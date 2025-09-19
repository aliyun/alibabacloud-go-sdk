// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemWarningMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckItemWarningMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckItemWarningMachineResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckItemWarningMachineResponseBody) *ListCheckItemWarningMachineResponse
	GetBody() *ListCheckItemWarningMachineResponseBody
}

type ListCheckItemWarningMachineResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckItemWarningMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckItemWarningMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineResponse) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckItemWarningMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckItemWarningMachineResponse) GetBody() *ListCheckItemWarningMachineResponseBody {
	return s.Body
}

func (s *ListCheckItemWarningMachineResponse) SetHeaders(v map[string]*string) *ListCheckItemWarningMachineResponse {
	s.Headers = v
	return s
}

func (s *ListCheckItemWarningMachineResponse) SetStatusCode(v int32) *ListCheckItemWarningMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckItemWarningMachineResponse) SetBody(v *ListCheckItemWarningMachineResponseBody) *ListCheckItemWarningMachineResponse {
	s.Body = v
	return s
}

func (s *ListCheckItemWarningMachineResponse) Validate() error {
	return dara.Validate(s)
}
