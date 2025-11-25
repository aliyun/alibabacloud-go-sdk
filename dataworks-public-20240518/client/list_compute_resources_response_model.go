// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeResourcesResponseBody) *ListComputeResourcesResponse
	GetBody() *ListComputeResourcesResponseBody
}

type ListComputeResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeResourcesResponse) GetBody() *ListComputeResourcesResponseBody {
	return s.Body
}

func (s *ListComputeResourcesResponse) SetHeaders(v map[string]*string) *ListComputeResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListComputeResourcesResponse) SetStatusCode(v int32) *ListComputeResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeResourcesResponse) SetBody(v *ListComputeResourcesResponseBody) *ListComputeResourcesResponse {
	s.Body = v
	return s
}

func (s *ListComputeResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
