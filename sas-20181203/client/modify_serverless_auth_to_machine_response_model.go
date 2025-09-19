// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyServerlessAuthToMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyServerlessAuthToMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyServerlessAuthToMachineResponse
	GetStatusCode() *int32
	SetBody(v *ModifyServerlessAuthToMachineResponseBody) *ModifyServerlessAuthToMachineResponse
	GetBody() *ModifyServerlessAuthToMachineResponseBody
}

type ModifyServerlessAuthToMachineResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyServerlessAuthToMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyServerlessAuthToMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyServerlessAuthToMachineResponse) GoString() string {
	return s.String()
}

func (s *ModifyServerlessAuthToMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyServerlessAuthToMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyServerlessAuthToMachineResponse) GetBody() *ModifyServerlessAuthToMachineResponseBody {
	return s.Body
}

func (s *ModifyServerlessAuthToMachineResponse) SetHeaders(v map[string]*string) *ModifyServerlessAuthToMachineResponse {
	s.Headers = v
	return s
}

func (s *ModifyServerlessAuthToMachineResponse) SetStatusCode(v int32) *ModifyServerlessAuthToMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyServerlessAuthToMachineResponse) SetBody(v *ModifyServerlessAuthToMachineResponseBody) *ModifyServerlessAuthToMachineResponse {
	s.Body = v
	return s
}

func (s *ModifyServerlessAuthToMachineResponse) Validate() error {
	return dara.Validate(s)
}
