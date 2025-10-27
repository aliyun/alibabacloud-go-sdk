// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateK8sResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateK8sResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateK8sResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateK8sResponseBody) *ListPrivateK8sResponse
	GetBody() *ListPrivateK8sResponseBody
}

type ListPrivateK8sResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateK8sResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateK8sResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateK8sResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateK8sResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateK8sResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateK8sResponse) GetBody() *ListPrivateK8sResponseBody {
	return s.Body
}

func (s *ListPrivateK8sResponse) SetHeaders(v map[string]*string) *ListPrivateK8sResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateK8sResponse) SetStatusCode(v int32) *ListPrivateK8sResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateK8sResponse) SetBody(v *ListPrivateK8sResponseBody) *ListPrivateK8sResponse {
	s.Body = v
	return s
}

func (s *ListPrivateK8sResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
