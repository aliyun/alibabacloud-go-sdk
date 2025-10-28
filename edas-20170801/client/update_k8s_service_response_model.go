// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sServiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sServiceResponseBody) *UpdateK8sServiceResponse
	GetBody() *UpdateK8sServiceResponseBody
}

type UpdateK8sServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sServiceResponse) GetBody() *UpdateK8sServiceResponseBody {
	return s.Body
}

func (s *UpdateK8sServiceResponse) SetHeaders(v map[string]*string) *UpdateK8sServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sServiceResponse) SetStatusCode(v int32) *UpdateK8sServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sServiceResponse) SetBody(v *UpdateK8sServiceResponseBody) *UpdateK8sServiceResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
