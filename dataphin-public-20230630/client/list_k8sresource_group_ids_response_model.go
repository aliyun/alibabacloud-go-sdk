// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8SResourceGroupIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8SResourceGroupIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8SResourceGroupIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListK8SResourceGroupIdsResponseBody) *ListK8SResourceGroupIdsResponse
	GetBody() *ListK8SResourceGroupIdsResponseBody
}

type ListK8SResourceGroupIdsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8SResourceGroupIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8SResourceGroupIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8SResourceGroupIdsResponse) GoString() string {
	return s.String()
}

func (s *ListK8SResourceGroupIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8SResourceGroupIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8SResourceGroupIdsResponse) GetBody() *ListK8SResourceGroupIdsResponseBody {
	return s.Body
}

func (s *ListK8SResourceGroupIdsResponse) SetHeaders(v map[string]*string) *ListK8SResourceGroupIdsResponse {
	s.Headers = v
	return s
}

func (s *ListK8SResourceGroupIdsResponse) SetStatusCode(v int32) *ListK8SResourceGroupIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8SResourceGroupIdsResponse) SetBody(v *ListK8SResourceGroupIdsResponseBody) *ListK8SResourceGroupIdsResponse {
	s.Body = v
	return s
}

func (s *ListK8SResourceGroupIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
