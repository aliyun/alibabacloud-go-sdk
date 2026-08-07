// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlashSmsAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFlashSmsAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFlashSmsAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *CreateFlashSmsAccessProfileResponseBody) *CreateFlashSmsAccessProfileResponse
	GetBody() *CreateFlashSmsAccessProfileResponseBody
}

type CreateFlashSmsAccessProfileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlashSmsAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFlashSmsAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFlashSmsAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateFlashSmsAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFlashSmsAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFlashSmsAccessProfileResponse) GetBody() *CreateFlashSmsAccessProfileResponseBody {
	return s.Body
}

func (s *CreateFlashSmsAccessProfileResponse) SetHeaders(v map[string]*string) *CreateFlashSmsAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateFlashSmsAccessProfileResponse) SetStatusCode(v int32) *CreateFlashSmsAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlashSmsAccessProfileResponse) SetBody(v *CreateFlashSmsAccessProfileResponseBody) *CreateFlashSmsAccessProfileResponse {
	s.Body = v
	return s
}

func (s *CreateFlashSmsAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
