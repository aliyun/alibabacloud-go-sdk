// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodePoolComponentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodePoolComponentsResponse
	GetStatusCode() *int32
	SetBody(v *ListNodePoolComponentsResponseBody) *ListNodePoolComponentsResponse
	GetBody() *ListNodePoolComponentsResponseBody
}

type ListNodePoolComponentsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodePoolComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodePoolComponentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodePoolComponentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodePoolComponentsResponse) GetBody() *ListNodePoolComponentsResponseBody {
	return s.Body
}

func (s *ListNodePoolComponentsResponse) SetHeaders(v map[string]*string) *ListNodePoolComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListNodePoolComponentsResponse) SetStatusCode(v int32) *ListNodePoolComponentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodePoolComponentsResponse) SetBody(v *ListNodePoolComponentsResponseBody) *ListNodePoolComponentsResponse {
	s.Body = v
	return s
}

func (s *ListNodePoolComponentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
