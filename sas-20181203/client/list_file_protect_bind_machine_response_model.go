// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectBindMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectBindMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectBindMachineResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectBindMachineResponseBody) *ListFileProtectBindMachineResponse
	GetBody() *ListFileProtectBindMachineResponseBody
}

type ListFileProtectBindMachineResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectBindMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectBindMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectBindMachineResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectBindMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectBindMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectBindMachineResponse) GetBody() *ListFileProtectBindMachineResponseBody {
	return s.Body
}

func (s *ListFileProtectBindMachineResponse) SetHeaders(v map[string]*string) *ListFileProtectBindMachineResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectBindMachineResponse) SetStatusCode(v int32) *ListFileProtectBindMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectBindMachineResponse) SetBody(v *ListFileProtectBindMachineResponseBody) *ListFileProtectBindMachineResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectBindMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
