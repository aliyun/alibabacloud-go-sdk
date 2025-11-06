// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMdsCubeResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMdsCubeResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListMdsCubeResourcesResponseBody) *ListMdsCubeResourcesResponse
	GetBody() *ListMdsCubeResourcesResponseBody
}

type ListMdsCubeResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMdsCubeResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMdsCubeResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListMdsCubeResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMdsCubeResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMdsCubeResourcesResponse) GetBody() *ListMdsCubeResourcesResponseBody {
	return s.Body
}

func (s *ListMdsCubeResourcesResponse) SetHeaders(v map[string]*string) *ListMdsCubeResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListMdsCubeResourcesResponse) SetStatusCode(v int32) *ListMdsCubeResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMdsCubeResourcesResponse) SetBody(v *ListMdsCubeResourcesResponseBody) *ListMdsCubeResourcesResponse {
	s.Body = v
	return s
}

func (s *ListMdsCubeResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
