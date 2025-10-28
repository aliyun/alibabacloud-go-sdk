// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sResourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sResourceResponseBody) *UpdateK8sResourceResponse
	GetBody() *UpdateK8sResourceResponseBody
}

type UpdateK8sResourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sResourceResponse) GetBody() *UpdateK8sResourceResponseBody {
	return s.Body
}

func (s *UpdateK8sResourceResponse) SetHeaders(v map[string]*string) *UpdateK8sResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sResourceResponse) SetStatusCode(v int32) *UpdateK8sResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sResourceResponse) SetBody(v *UpdateK8sResourceResponseBody) *UpdateK8sResourceResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
