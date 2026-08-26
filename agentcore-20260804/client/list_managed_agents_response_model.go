// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListManagedAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListManagedAgentsResponse
	GetStatusCode() *int32
	SetBody(v *ListManagedAgentsResponseBody) *ListManagedAgentsResponse
	GetBody() *ListManagedAgentsResponseBody
}

type ListManagedAgentsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListManagedAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListManagedAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListManagedAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListManagedAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListManagedAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListManagedAgentsResponse) GetBody() *ListManagedAgentsResponseBody {
	return s.Body
}

func (s *ListManagedAgentsResponse) SetHeaders(v map[string]*string) *ListManagedAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListManagedAgentsResponse) SetStatusCode(v int32) *ListManagedAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagedAgentsResponse) SetBody(v *ListManagedAgentsResponseBody) *ListManagedAgentsResponse {
	s.Body = v
	return s
}

func (s *ListManagedAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
