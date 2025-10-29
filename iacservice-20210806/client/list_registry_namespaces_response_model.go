// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRegistryNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRegistryNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *ListRegistryNamespacesResponseBody) *ListRegistryNamespacesResponse
	GetBody() *ListRegistryNamespacesResponseBody
}

type ListRegistryNamespacesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistryNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistryNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryNamespacesResponse) GoString() string {
	return s.String()
}

func (s *ListRegistryNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRegistryNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRegistryNamespacesResponse) GetBody() *ListRegistryNamespacesResponseBody {
	return s.Body
}

func (s *ListRegistryNamespacesResponse) SetHeaders(v map[string]*string) *ListRegistryNamespacesResponse {
	s.Headers = v
	return s
}

func (s *ListRegistryNamespacesResponse) SetStatusCode(v int32) *ListRegistryNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistryNamespacesResponse) SetBody(v *ListRegistryNamespacesResponseBody) *ListRegistryNamespacesResponse {
	s.Body = v
	return s
}

func (s *ListRegistryNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
