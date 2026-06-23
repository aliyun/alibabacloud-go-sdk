// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceServerlessSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceServerlessSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceServerlessSwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceServerlessSwitchResponseBody) *UpdateInstanceServerlessSwitchResponse
	GetBody() *UpdateInstanceServerlessSwitchResponseBody
}

type UpdateInstanceServerlessSwitchResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceServerlessSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceServerlessSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceServerlessSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceServerlessSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceServerlessSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceServerlessSwitchResponse) GetBody() *UpdateInstanceServerlessSwitchResponseBody {
	return s.Body
}

func (s *UpdateInstanceServerlessSwitchResponse) SetHeaders(v map[string]*string) *UpdateInstanceServerlessSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponse) SetStatusCode(v int32) *UpdateInstanceServerlessSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponse) SetBody(v *UpdateInstanceServerlessSwitchResponseBody) *UpdateInstanceServerlessSwitchResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceServerlessSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
