// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceNodesResponseBody) *ListInstanceNodesResponse
	GetBody() *ListInstanceNodesResponseBody
}

type ListInstanceNodesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceNodesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceNodesResponse) GetBody() *ListInstanceNodesResponseBody {
	return s.Body
}

func (s *ListInstanceNodesResponse) SetHeaders(v map[string]*string) *ListInstanceNodesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceNodesResponse) SetStatusCode(v int32) *ListInstanceNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceNodesResponse) SetBody(v *ListInstanceNodesResponseBody) *ListInstanceNodesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
