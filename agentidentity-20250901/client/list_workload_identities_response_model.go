// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkloadIdentitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkloadIdentitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkloadIdentitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkloadIdentitiesResponseBody) *ListWorkloadIdentitiesResponse
	GetBody() *ListWorkloadIdentitiesResponseBody
}

type ListWorkloadIdentitiesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkloadIdentitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkloadIdentitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkloadIdentitiesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkloadIdentitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkloadIdentitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkloadIdentitiesResponse) GetBody() *ListWorkloadIdentitiesResponseBody {
	return s.Body
}

func (s *ListWorkloadIdentitiesResponse) SetHeaders(v map[string]*string) *ListWorkloadIdentitiesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkloadIdentitiesResponse) SetStatusCode(v int32) *ListWorkloadIdentitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkloadIdentitiesResponse) SetBody(v *ListWorkloadIdentitiesResponseBody) *ListWorkloadIdentitiesResponse {
	s.Body = v
	return s
}

func (s *ListWorkloadIdentitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
