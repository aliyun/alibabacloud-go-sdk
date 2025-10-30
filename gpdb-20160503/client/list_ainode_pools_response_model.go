// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAINodePoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAINodePoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAINodePoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListAINodePoolsResponseBody) *ListAINodePoolsResponse
	GetBody() *ListAINodePoolsResponseBody
}

type ListAINodePoolsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAINodePoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAINodePoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAINodePoolsResponse) GoString() string {
	return s.String()
}

func (s *ListAINodePoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAINodePoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAINodePoolsResponse) GetBody() *ListAINodePoolsResponseBody {
	return s.Body
}

func (s *ListAINodePoolsResponse) SetHeaders(v map[string]*string) *ListAINodePoolsResponse {
	s.Headers = v
	return s
}

func (s *ListAINodePoolsResponse) SetStatusCode(v int32) *ListAINodePoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAINodePoolsResponse) SetBody(v *ListAINodePoolsResponseBody) *ListAINodePoolsResponse {
	s.Body = v
	return s
}

func (s *ListAINodePoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
