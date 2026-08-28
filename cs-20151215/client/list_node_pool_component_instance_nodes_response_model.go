// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePoolComponentInstanceNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodePoolComponentInstanceNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodePoolComponentInstanceNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListNodePoolComponentInstanceNodesResponseBody) *ListNodePoolComponentInstanceNodesResponse
	GetBody() *ListNodePoolComponentInstanceNodesResponseBody
}

type ListNodePoolComponentInstanceNodesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodePoolComponentInstanceNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodePoolComponentInstanceNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodePoolComponentInstanceNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodePoolComponentInstanceNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodePoolComponentInstanceNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodePoolComponentInstanceNodesResponse) GetBody() *ListNodePoolComponentInstanceNodesResponseBody {
	return s.Body
}

func (s *ListNodePoolComponentInstanceNodesResponse) SetHeaders(v map[string]*string) *ListNodePoolComponentInstanceNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponse) SetStatusCode(v int32) *ListNodePoolComponentInstanceNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponse) SetBody(v *ListNodePoolComponentInstanceNodesResponseBody) *ListNodePoolComponentInstanceNodesResponse {
	s.Body = v
	return s
}

func (s *ListNodePoolComponentInstanceNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
