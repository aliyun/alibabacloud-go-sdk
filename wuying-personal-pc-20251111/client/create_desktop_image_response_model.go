// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDesktopImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDesktopImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateDesktopImageResponseBody) *CreateDesktopImageResponse
	GetBody() *CreateDesktopImageResponseBody
}

type CreateDesktopImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDesktopImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDesktopImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopImageResponse) GoString() string {
	return s.String()
}

func (s *CreateDesktopImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDesktopImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDesktopImageResponse) GetBody() *CreateDesktopImageResponseBody {
	return s.Body
}

func (s *CreateDesktopImageResponse) SetHeaders(v map[string]*string) *CreateDesktopImageResponse {
	s.Headers = v
	return s
}

func (s *CreateDesktopImageResponse) SetStatusCode(v int32) *CreateDesktopImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDesktopImageResponse) SetBody(v *CreateDesktopImageResponseBody) *CreateDesktopImageResponse {
	s.Body = v
	return s
}

func (s *CreateDesktopImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
