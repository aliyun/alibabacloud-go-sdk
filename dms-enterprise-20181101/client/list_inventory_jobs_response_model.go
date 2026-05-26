// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventoryJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInventoryJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInventoryJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListInventoryJobsResponseBody) *ListInventoryJobsResponse
	GetBody() *ListInventoryJobsResponseBody
}

type ListInventoryJobsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInventoryJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInventoryJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryJobsResponse) GoString() string {
	return s.String()
}

func (s *ListInventoryJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInventoryJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInventoryJobsResponse) GetBody() *ListInventoryJobsResponseBody {
	return s.Body
}

func (s *ListInventoryJobsResponse) SetHeaders(v map[string]*string) *ListInventoryJobsResponse {
	s.Headers = v
	return s
}

func (s *ListInventoryJobsResponse) SetStatusCode(v int32) *ListInventoryJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInventoryJobsResponse) SetBody(v *ListInventoryJobsResponseBody) *ListInventoryJobsResponse {
	s.Body = v
	return s
}

func (s *ListInventoryJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
