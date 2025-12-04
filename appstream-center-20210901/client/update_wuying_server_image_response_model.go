// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWuyingServerImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWuyingServerImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWuyingServerImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWuyingServerImageResponseBody) *UpdateWuyingServerImageResponse
	GetBody() *UpdateWuyingServerImageResponseBody
}

type UpdateWuyingServerImageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWuyingServerImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWuyingServerImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWuyingServerImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateWuyingServerImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWuyingServerImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWuyingServerImageResponse) GetBody() *UpdateWuyingServerImageResponseBody {
	return s.Body
}

func (s *UpdateWuyingServerImageResponse) SetHeaders(v map[string]*string) *UpdateWuyingServerImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateWuyingServerImageResponse) SetStatusCode(v int32) *UpdateWuyingServerImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWuyingServerImageResponse) SetBody(v *UpdateWuyingServerImageResponseBody) *UpdateWuyingServerImageResponse {
	s.Body = v
	return s
}

func (s *UpdateWuyingServerImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
