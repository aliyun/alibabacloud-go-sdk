// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsCodeLimitConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSmsCodeLimitConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSmsCodeLimitConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSmsCodeLimitConfigResponseBody) *UpdateSmsCodeLimitConfigResponse
	GetBody() *UpdateSmsCodeLimitConfigResponseBody
}

type UpdateSmsCodeLimitConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSmsCodeLimitConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSmsCodeLimitConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsCodeLimitConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmsCodeLimitConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSmsCodeLimitConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSmsCodeLimitConfigResponse) GetBody() *UpdateSmsCodeLimitConfigResponseBody {
	return s.Body
}

func (s *UpdateSmsCodeLimitConfigResponse) SetHeaders(v map[string]*string) *UpdateSmsCodeLimitConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponse) SetStatusCode(v int32) *UpdateSmsCodeLimitConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponse) SetBody(v *UpdateSmsCodeLimitConfigResponseBody) *UpdateSmsCodeLimitConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateSmsCodeLimitConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
