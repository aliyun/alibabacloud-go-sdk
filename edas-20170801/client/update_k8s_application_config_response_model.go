// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sApplicationConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sApplicationConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sApplicationConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sApplicationConfigResponseBody) *UpdateK8sApplicationConfigResponse
	GetBody() *UpdateK8sApplicationConfigResponseBody
}

type UpdateK8sApplicationConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sApplicationConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sApplicationConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sApplicationConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sApplicationConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sApplicationConfigResponse) GetBody() *UpdateK8sApplicationConfigResponseBody {
	return s.Body
}

func (s *UpdateK8sApplicationConfigResponse) SetHeaders(v map[string]*string) *UpdateK8sApplicationConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sApplicationConfigResponse) SetStatusCode(v int32) *UpdateK8sApplicationConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sApplicationConfigResponse) SetBody(v *UpdateK8sApplicationConfigResponseBody) *UpdateK8sApplicationConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sApplicationConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
