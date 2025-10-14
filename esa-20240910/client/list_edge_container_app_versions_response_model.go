// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeContainerAppVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeContainerAppVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeContainerAppVersionsResponseBody) *ListEdgeContainerAppVersionsResponse
	GetBody() *ListEdgeContainerAppVersionsResponseBody
}

type ListEdgeContainerAppVersionsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeContainerAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeContainerAppVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeContainerAppVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeContainerAppVersionsResponse) GetBody() *ListEdgeContainerAppVersionsResponseBody {
	return s.Body
}

func (s *ListEdgeContainerAppVersionsResponse) SetHeaders(v map[string]*string) *ListEdgeContainerAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeContainerAppVersionsResponse) SetStatusCode(v int32) *ListEdgeContainerAppVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeContainerAppVersionsResponse) SetBody(v *ListEdgeContainerAppVersionsResponseBody) *ListEdgeContainerAppVersionsResponse {
	s.Body = v
	return s
}

func (s *ListEdgeContainerAppVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
