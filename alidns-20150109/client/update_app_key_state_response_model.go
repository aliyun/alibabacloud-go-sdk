// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppKeyStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppKeyStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppKeyStateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppKeyStateResponseBody) *UpdateAppKeyStateResponse
	GetBody() *UpdateAppKeyStateResponseBody
}

type UpdateAppKeyStateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppKeyStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppKeyStateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppKeyStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppKeyStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppKeyStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppKeyStateResponse) GetBody() *UpdateAppKeyStateResponseBody {
	return s.Body
}

func (s *UpdateAppKeyStateResponse) SetHeaders(v map[string]*string) *UpdateAppKeyStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppKeyStateResponse) SetStatusCode(v int32) *UpdateAppKeyStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppKeyStateResponse) SetBody(v *UpdateAppKeyStateResponseBody) *UpdateAppKeyStateResponse {
	s.Body = v
	return s
}

func (s *UpdateAppKeyStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
