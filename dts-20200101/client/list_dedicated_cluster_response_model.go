// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDedicatedClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDedicatedClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDedicatedClusterResponse
	GetStatusCode() *int32
	SetBody(v *ListDedicatedClusterResponseBody) *ListDedicatedClusterResponse
	GetBody() *ListDedicatedClusterResponseBody
}

type ListDedicatedClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDedicatedClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDedicatedClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDedicatedClusterResponse) GoString() string {
	return s.String()
}

func (s *ListDedicatedClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDedicatedClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDedicatedClusterResponse) GetBody() *ListDedicatedClusterResponseBody {
	return s.Body
}

func (s *ListDedicatedClusterResponse) SetHeaders(v map[string]*string) *ListDedicatedClusterResponse {
	s.Headers = v
	return s
}

func (s *ListDedicatedClusterResponse) SetStatusCode(v int32) *ListDedicatedClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDedicatedClusterResponse) SetBody(v *ListDedicatedClusterResponseBody) *ListDedicatedClusterResponse {
	s.Body = v
	return s
}

func (s *ListDedicatedClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
