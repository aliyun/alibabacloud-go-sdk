// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCImageResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCImageResponseBody) *CreateRCImageResponse
	GetBody() *CreateRCImageResponseBody
}

type CreateRCImageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCImageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCImageResponse) GoString() string {
	return s.String()
}

func (s *CreateRCImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCImageResponse) GetBody() *CreateRCImageResponseBody {
	return s.Body
}

func (s *CreateRCImageResponse) SetHeaders(v map[string]*string) *CreateRCImageResponse {
	s.Headers = v
	return s
}

func (s *CreateRCImageResponse) SetStatusCode(v int32) *CreateRCImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCImageResponse) SetBody(v *CreateRCImageResponseBody) *CreateRCImageResponse {
	s.Body = v
	return s
}

func (s *CreateRCImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
