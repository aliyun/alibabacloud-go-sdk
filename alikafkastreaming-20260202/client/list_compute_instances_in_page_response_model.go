// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesInPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeInstancesInPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeInstancesInPageResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeInstancesInPageResponseBody) *ListComputeInstancesInPageResponse
	GetBody() *ListComputeInstancesInPageResponseBody
}

type ListComputeInstancesInPageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeInstancesInPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeInstancesInPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesInPageResponse) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesInPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeInstancesInPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeInstancesInPageResponse) GetBody() *ListComputeInstancesInPageResponseBody {
	return s.Body
}

func (s *ListComputeInstancesInPageResponse) SetHeaders(v map[string]*string) *ListComputeInstancesInPageResponse {
	s.Headers = v
	return s
}

func (s *ListComputeInstancesInPageResponse) SetStatusCode(v int32) *ListComputeInstancesInPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeInstancesInPageResponse) SetBody(v *ListComputeInstancesInPageResponseBody) *ListComputeInstancesInPageResponse {
	s.Body = v
	return s
}

func (s *ListComputeInstancesInPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
