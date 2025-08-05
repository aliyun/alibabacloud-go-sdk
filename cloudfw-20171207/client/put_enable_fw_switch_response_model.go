// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEnableFwSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutEnableFwSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutEnableFwSwitchResponse
	GetStatusCode() *int32
	SetBody(v *PutEnableFwSwitchResponseBody) *PutEnableFwSwitchResponse
	GetBody() *PutEnableFwSwitchResponseBody
}

type PutEnableFwSwitchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEnableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEnableFwSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s PutEnableFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutEnableFwSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutEnableFwSwitchResponse) GetBody() *PutEnableFwSwitchResponseBody {
	return s.Body
}

func (s *PutEnableFwSwitchResponse) SetHeaders(v map[string]*string) *PutEnableFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutEnableFwSwitchResponse) SetStatusCode(v int32) *PutEnableFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnableFwSwitchResponse) SetBody(v *PutEnableFwSwitchResponseBody) *PutEnableFwSwitchResponse {
	s.Body = v
	return s
}

func (s *PutEnableFwSwitchResponse) Validate() error {
	return dara.Validate(s)
}
