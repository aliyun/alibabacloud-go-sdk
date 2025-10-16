// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDisableFwSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDisableFwSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDisableFwSwitchResponse
	GetStatusCode() *int32
	SetBody(v *PutDisableFwSwitchResponseBody) *PutDisableFwSwitchResponse
	GetBody() *PutDisableFwSwitchResponseBody
}

type PutDisableFwSwitchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDisableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDisableFwSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDisableFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDisableFwSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDisableFwSwitchResponse) GetBody() *PutDisableFwSwitchResponseBody {
	return s.Body
}

func (s *PutDisableFwSwitchResponse) SetHeaders(v map[string]*string) *PutDisableFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutDisableFwSwitchResponse) SetStatusCode(v int32) *PutDisableFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDisableFwSwitchResponse) SetBody(v *PutDisableFwSwitchResponseBody) *PutDisableFwSwitchResponse {
	s.Body = v
	return s
}

func (s *PutDisableFwSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
