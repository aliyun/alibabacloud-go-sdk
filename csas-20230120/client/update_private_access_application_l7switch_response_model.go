// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationL7SwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateAccessApplicationL7SwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateAccessApplicationL7SwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateAccessApplicationL7SwitchResponseBody) *UpdatePrivateAccessApplicationL7SwitchResponse
	GetBody() *UpdatePrivateAccessApplicationL7SwitchResponseBody
}

type UpdatePrivateAccessApplicationL7SwitchResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateAccessApplicationL7SwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateAccessApplicationL7SwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationL7SwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) GetBody() *UpdatePrivateAccessApplicationL7SwitchResponseBody {
	return s.Body
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) SetHeaders(v map[string]*string) *UpdatePrivateAccessApplicationL7SwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) SetStatusCode(v int32) *UpdatePrivateAccessApplicationL7SwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) SetBody(v *UpdatePrivateAccessApplicationL7SwitchResponseBody) *UpdatePrivateAccessApplicationL7SwitchResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateAccessApplicationL7SwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
