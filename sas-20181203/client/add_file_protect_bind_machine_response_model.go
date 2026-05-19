// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileProtectBindMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFileProtectBindMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFileProtectBindMachineResponse
	GetStatusCode() *int32
	SetBody(v *AddFileProtectBindMachineResponseBody) *AddFileProtectBindMachineResponse
	GetBody() *AddFileProtectBindMachineResponseBody
}

type AddFileProtectBindMachineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFileProtectBindMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFileProtectBindMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFileProtectBindMachineResponse) GoString() string {
	return s.String()
}

func (s *AddFileProtectBindMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFileProtectBindMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFileProtectBindMachineResponse) GetBody() *AddFileProtectBindMachineResponseBody {
	return s.Body
}

func (s *AddFileProtectBindMachineResponse) SetHeaders(v map[string]*string) *AddFileProtectBindMachineResponse {
	s.Headers = v
	return s
}

func (s *AddFileProtectBindMachineResponse) SetStatusCode(v int32) *AddFileProtectBindMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFileProtectBindMachineResponse) SetBody(v *AddFileProtectBindMachineResponseBody) *AddFileProtectBindMachineResponse {
	s.Body = v
	return s
}

func (s *AddFileProtectBindMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
