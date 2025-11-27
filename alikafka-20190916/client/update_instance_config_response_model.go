// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceConfigResponseBody) *UpdateInstanceConfigResponse
	GetBody() *UpdateInstanceConfigResponseBody
}

type UpdateInstanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceConfigResponse) GetBody() *UpdateInstanceConfigResponseBody {
	return s.Body
}

func (s *UpdateInstanceConfigResponse) SetHeaders(v map[string]*string) *UpdateInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceConfigResponse) SetStatusCode(v int32) *UpdateInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceConfigResponse) SetBody(v *UpdateInstanceConfigResponseBody) *UpdateInstanceConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
