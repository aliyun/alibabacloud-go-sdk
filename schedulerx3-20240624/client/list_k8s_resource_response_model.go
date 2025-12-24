// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sResourceResponseBody) *ListK8sResourceResponse
	GetBody() *ListK8sResourceResponseBody
}

type ListK8sResourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sResourceResponse) GoString() string {
	return s.String()
}

func (s *ListK8sResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sResourceResponse) GetBody() *ListK8sResourceResponseBody {
	return s.Body
}

func (s *ListK8sResourceResponse) SetHeaders(v map[string]*string) *ListK8sResourceResponse {
	s.Headers = v
	return s
}

func (s *ListK8sResourceResponse) SetStatusCode(v int32) *ListK8sResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sResourceResponse) SetBody(v *ListK8sResourceResponseBody) *ListK8sResourceResponse {
	s.Body = v
	return s
}

func (s *ListK8sResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
