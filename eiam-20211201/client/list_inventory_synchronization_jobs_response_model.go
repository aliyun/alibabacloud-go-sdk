// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventorySynchronizationJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInventorySynchronizationJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInventorySynchronizationJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListInventorySynchronizationJobsResponseBody) *ListInventorySynchronizationJobsResponse
	GetBody() *ListInventorySynchronizationJobsResponseBody
}

type ListInventorySynchronizationJobsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInventorySynchronizationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInventorySynchronizationJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponse) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInventorySynchronizationJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInventorySynchronizationJobsResponse) GetBody() *ListInventorySynchronizationJobsResponseBody {
	return s.Body
}

func (s *ListInventorySynchronizationJobsResponse) SetHeaders(v map[string]*string) *ListInventorySynchronizationJobsResponse {
	s.Headers = v
	return s
}

func (s *ListInventorySynchronizationJobsResponse) SetStatusCode(v int32) *ListInventorySynchronizationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponse) SetBody(v *ListInventorySynchronizationJobsResponseBody) *ListInventorySynchronizationJobsResponse {
	s.Body = v
	return s
}

func (s *ListInventorySynchronizationJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
