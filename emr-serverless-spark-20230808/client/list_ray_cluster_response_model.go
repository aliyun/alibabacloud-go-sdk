// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRayClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRayClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRayClusterResponse
	GetStatusCode() *int32
	SetBody(v *ListRayClusterResponseBody) *ListRayClusterResponse
	GetBody() *ListRayClusterResponseBody
}

type ListRayClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRayClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRayClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRayClusterResponse) GoString() string {
	return s.String()
}

func (s *ListRayClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRayClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRayClusterResponse) GetBody() *ListRayClusterResponseBody {
	return s.Body
}

func (s *ListRayClusterResponse) SetHeaders(v map[string]*string) *ListRayClusterResponse {
	s.Headers = v
	return s
}

func (s *ListRayClusterResponse) SetStatusCode(v int32) *ListRayClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRayClusterResponse) SetBody(v *ListRayClusterResponseBody) *ListRayClusterResponse {
	s.Body = v
	return s
}

func (s *ListRayClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
