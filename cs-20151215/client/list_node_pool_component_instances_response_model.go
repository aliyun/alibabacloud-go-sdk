// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodePoolComponentInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodePoolComponentInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListNodePoolComponentInstancesResponseBody) *ListNodePoolComponentInstancesResponse
	GetBody() *ListNodePoolComponentInstancesResponseBody
}

type ListNodePoolComponentInstancesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodePoolComponentInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodePoolComponentInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodePoolComponentInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodePoolComponentInstancesResponse) GetBody() *ListNodePoolComponentInstancesResponseBody {
	return s.Body
}

func (s *ListNodePoolComponentInstancesResponse) SetHeaders(v map[string]*string) *ListNodePoolComponentInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListNodePoolComponentInstancesResponse) SetStatusCode(v int32) *ListNodePoolComponentInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodePoolComponentInstancesResponse) SetBody(v *ListNodePoolComponentInstancesResponseBody) *ListNodePoolComponentInstancesResponse {
	s.Body = v
	return s
}

func (s *ListNodePoolComponentInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
