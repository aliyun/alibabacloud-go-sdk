// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGuestClusterConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGuestClusterConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGuestClusterConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGuestClusterConfigResponseBody) *UpdateGuestClusterConfigResponse
	GetBody() *UpdateGuestClusterConfigResponseBody
}

type UpdateGuestClusterConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGuestClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGuestClusterConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGuestClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateGuestClusterConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGuestClusterConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGuestClusterConfigResponse) GetBody() *UpdateGuestClusterConfigResponseBody {
	return s.Body
}

func (s *UpdateGuestClusterConfigResponse) SetHeaders(v map[string]*string) *UpdateGuestClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateGuestClusterConfigResponse) SetStatusCode(v int32) *UpdateGuestClusterConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGuestClusterConfigResponse) SetBody(v *UpdateGuestClusterConfigResponseBody) *UpdateGuestClusterConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateGuestClusterConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
