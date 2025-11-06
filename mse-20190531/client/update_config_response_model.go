// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigResponseBody) *UpdateConfigResponse
	GetBody() *UpdateConfigResponseBody
}

type UpdateConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigResponse) GetBody() *UpdateConfigResponseBody {
	return s.Body
}

func (s *UpdateConfigResponse) SetHeaders(v map[string]*string) *UpdateConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigResponse) SetStatusCode(v int32) *UpdateConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigResponse) SetBody(v *UpdateConfigResponseBody) *UpdateConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
