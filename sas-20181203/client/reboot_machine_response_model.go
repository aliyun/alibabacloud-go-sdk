// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootMachineResponse
	GetStatusCode() *int32
	SetBody(v *RebootMachineResponseBody) *RebootMachineResponse
	GetBody() *RebootMachineResponseBody
}

type RebootMachineResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootMachineResponse) GoString() string {
	return s.String()
}

func (s *RebootMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootMachineResponse) GetBody() *RebootMachineResponseBody {
	return s.Body
}

func (s *RebootMachineResponse) SetHeaders(v map[string]*string) *RebootMachineResponse {
	s.Headers = v
	return s
}

func (s *RebootMachineResponse) SetStatusCode(v int32) *RebootMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootMachineResponse) SetBody(v *RebootMachineResponseBody) *RebootMachineResponse {
	s.Body = v
	return s
}

func (s *RebootMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
