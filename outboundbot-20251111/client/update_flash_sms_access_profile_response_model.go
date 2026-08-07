// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFlashSmsAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFlashSmsAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFlashSmsAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFlashSmsAccessProfileResponseBody) *UpdateFlashSmsAccessProfileResponse
	GetBody() *UpdateFlashSmsAccessProfileResponseBody
}

type UpdateFlashSmsAccessProfileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlashSmsAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFlashSmsAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFlashSmsAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlashSmsAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFlashSmsAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFlashSmsAccessProfileResponse) GetBody() *UpdateFlashSmsAccessProfileResponseBody {
	return s.Body
}

func (s *UpdateFlashSmsAccessProfileResponse) SetHeaders(v map[string]*string) *UpdateFlashSmsAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponse) SetStatusCode(v int32) *UpdateFlashSmsAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponse) SetBody(v *UpdateFlashSmsAccessProfileResponseBody) *UpdateFlashSmsAccessProfileResponse {
	s.Body = v
	return s
}

func (s *UpdateFlashSmsAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
