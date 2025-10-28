// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sConfigMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sConfigMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sConfigMapResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sConfigMapResponseBody) *UpdateK8sConfigMapResponse
	GetBody() *UpdateK8sConfigMapResponseBody
}

type UpdateK8sConfigMapResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sConfigMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sConfigMapResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sConfigMapResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sConfigMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sConfigMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sConfigMapResponse) GetBody() *UpdateK8sConfigMapResponseBody {
	return s.Body
}

func (s *UpdateK8sConfigMapResponse) SetHeaders(v map[string]*string) *UpdateK8sConfigMapResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sConfigMapResponse) SetStatusCode(v int32) *UpdateK8sConfigMapResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sConfigMapResponse) SetBody(v *UpdateK8sConfigMapResponseBody) *UpdateK8sConfigMapResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sConfigMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
