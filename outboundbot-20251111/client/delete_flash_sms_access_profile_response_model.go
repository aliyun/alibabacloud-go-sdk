// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlashSmsAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFlashSmsAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFlashSmsAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFlashSmsAccessProfileResponseBody) *DeleteFlashSmsAccessProfileResponse
	GetBody() *DeleteFlashSmsAccessProfileResponseBody
}

type DeleteFlashSmsAccessProfileResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlashSmsAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFlashSmsAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlashSmsAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlashSmsAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFlashSmsAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFlashSmsAccessProfileResponse) GetBody() *DeleteFlashSmsAccessProfileResponseBody {
	return s.Body
}

func (s *DeleteFlashSmsAccessProfileResponse) SetHeaders(v map[string]*string) *DeleteFlashSmsAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponse) SetStatusCode(v int32) *DeleteFlashSmsAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponse) SetBody(v *DeleteFlashSmsAccessProfileResponseBody) *DeleteFlashSmsAccessProfileResponse {
	s.Body = v
	return s
}

func (s *DeleteFlashSmsAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
