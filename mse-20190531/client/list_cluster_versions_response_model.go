// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterVersionsResponseBody) *ListClusterVersionsResponse
	GetBody() *ListClusterVersionsResponseBody
}

type ListClusterVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterVersionsResponse) GetBody() *ListClusterVersionsResponseBody {
	return s.Body
}

func (s *ListClusterVersionsResponse) SetHeaders(v map[string]*string) *ListClusterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterVersionsResponse) SetStatusCode(v int32) *ListClusterVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterVersionsResponse) SetBody(v *ListClusterVersionsResponseBody) *ListClusterVersionsResponse {
	s.Body = v
	return s
}

func (s *ListClusterVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
