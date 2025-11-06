// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerlessSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServerlessSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServerlessSwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServerlessSwitchResponseBody) *UpdateServerlessSwitchResponse
	GetBody() *UpdateServerlessSwitchResponseBody
}

type UpdateServerlessSwitchResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServerlessSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServerlessSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerlessSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateServerlessSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServerlessSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServerlessSwitchResponse) GetBody() *UpdateServerlessSwitchResponseBody {
	return s.Body
}

func (s *UpdateServerlessSwitchResponse) SetHeaders(v map[string]*string) *UpdateServerlessSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateServerlessSwitchResponse) SetStatusCode(v int32) *UpdateServerlessSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServerlessSwitchResponse) SetBody(v *UpdateServerlessSwitchResponseBody) *UpdateServerlessSwitchResponse {
	s.Body = v
	return s
}

func (s *UpdateServerlessSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
