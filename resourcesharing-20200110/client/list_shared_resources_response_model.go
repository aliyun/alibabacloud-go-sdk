// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSharedResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSharedResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSharedResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListSharedResourcesResponseBody) *ListSharedResourcesResponse
	GetBody() *ListSharedResourcesResponseBody
}

type ListSharedResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSharedResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSharedResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSharedResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListSharedResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSharedResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSharedResourcesResponse) GetBody() *ListSharedResourcesResponseBody {
	return s.Body
}

func (s *ListSharedResourcesResponse) SetHeaders(v map[string]*string) *ListSharedResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListSharedResourcesResponse) SetStatusCode(v int32) *ListSharedResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSharedResourcesResponse) SetBody(v *ListSharedResourcesResponseBody) *ListSharedResourcesResponse {
	s.Body = v
	return s
}

func (s *ListSharedResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
