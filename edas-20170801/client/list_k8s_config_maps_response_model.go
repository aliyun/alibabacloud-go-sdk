// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sConfigMapsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sConfigMapsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sConfigMapsResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sConfigMapsResponseBody) *ListK8sConfigMapsResponse
	GetBody() *ListK8sConfigMapsResponseBody
}

type ListK8sConfigMapsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sConfigMapsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sConfigMapsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sConfigMapsResponse) GoString() string {
	return s.String()
}

func (s *ListK8sConfigMapsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sConfigMapsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sConfigMapsResponse) GetBody() *ListK8sConfigMapsResponseBody {
	return s.Body
}

func (s *ListK8sConfigMapsResponse) SetHeaders(v map[string]*string) *ListK8sConfigMapsResponse {
	s.Headers = v
	return s
}

func (s *ListK8sConfigMapsResponse) SetStatusCode(v int32) *ListK8sConfigMapsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sConfigMapsResponse) SetBody(v *ListK8sConfigMapsResponseBody) *ListK8sConfigMapsResponse {
	s.Body = v
	return s
}

func (s *ListK8sConfigMapsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
