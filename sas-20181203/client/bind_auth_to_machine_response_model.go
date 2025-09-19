// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAuthToMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAuthToMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAuthToMachineResponse
	GetStatusCode() *int32
	SetBody(v *BindAuthToMachineResponseBody) *BindAuthToMachineResponse
	GetBody() *BindAuthToMachineResponseBody
}

type BindAuthToMachineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAuthToMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAuthToMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAuthToMachineResponse) GoString() string {
	return s.String()
}

func (s *BindAuthToMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAuthToMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAuthToMachineResponse) GetBody() *BindAuthToMachineResponseBody {
	return s.Body
}

func (s *BindAuthToMachineResponse) SetHeaders(v map[string]*string) *BindAuthToMachineResponse {
	s.Headers = v
	return s
}

func (s *BindAuthToMachineResponse) SetStatusCode(v int32) *BindAuthToMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAuthToMachineResponse) SetBody(v *BindAuthToMachineResponseBody) *BindAuthToMachineResponse {
	s.Body = v
	return s
}

func (s *BindAuthToMachineResponse) Validate() error {
	return dara.Validate(s)
}
