// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVsPullStreamInfoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVsPullStreamInfoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVsPullStreamInfoConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVsPullStreamInfoConfigResponseBody) *UpdateVsPullStreamInfoConfigResponse
	GetBody() *UpdateVsPullStreamInfoConfigResponseBody
}

type UpdateVsPullStreamInfoConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVsPullStreamInfoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVsPullStreamInfoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVsPullStreamInfoConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateVsPullStreamInfoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVsPullStreamInfoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVsPullStreamInfoConfigResponse) GetBody() *UpdateVsPullStreamInfoConfigResponseBody {
	return s.Body
}

func (s *UpdateVsPullStreamInfoConfigResponse) SetHeaders(v map[string]*string) *UpdateVsPullStreamInfoConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateVsPullStreamInfoConfigResponse) SetStatusCode(v int32) *UpdateVsPullStreamInfoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigResponse) SetBody(v *UpdateVsPullStreamInfoConfigResponseBody) *UpdateVsPullStreamInfoConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateVsPullStreamInfoConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
