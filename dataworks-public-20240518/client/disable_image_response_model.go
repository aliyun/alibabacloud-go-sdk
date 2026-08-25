// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableImageResponse
	GetStatusCode() *int32
	SetBody(v *DisableImageResponseBody) *DisableImageResponse
	GetBody() *DisableImageResponseBody
}

type DisableImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableImageResponse) GoString() string {
	return s.String()
}

func (s *DisableImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableImageResponse) GetBody() *DisableImageResponseBody {
	return s.Body
}

func (s *DisableImageResponse) SetHeaders(v map[string]*string) *DisableImageResponse {
	s.Headers = v
	return s
}

func (s *DisableImageResponse) SetStatusCode(v int32) *DisableImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableImageResponse) SetBody(v *DisableImageResponseBody) *DisableImageResponse {
	s.Body = v
	return s
}

func (s *DisableImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
