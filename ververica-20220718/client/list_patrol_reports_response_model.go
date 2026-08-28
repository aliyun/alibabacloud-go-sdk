// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatrolReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPatrolReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPatrolReportsResponse
	GetStatusCode() *int32
	SetBody(v *ListPatrolReportsResponseBody) *ListPatrolReportsResponse
	GetBody() *ListPatrolReportsResponseBody
}

type ListPatrolReportsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPatrolReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPatrolReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPatrolReportsResponse) GoString() string {
	return s.String()
}

func (s *ListPatrolReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPatrolReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPatrolReportsResponse) GetBody() *ListPatrolReportsResponseBody {
	return s.Body
}

func (s *ListPatrolReportsResponse) SetHeaders(v map[string]*string) *ListPatrolReportsResponse {
	s.Headers = v
	return s
}

func (s *ListPatrolReportsResponse) SetStatusCode(v int32) *ListPatrolReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPatrolReportsResponse) SetBody(v *ListPatrolReportsResponseBody) *ListPatrolReportsResponse {
	s.Body = v
	return s
}

func (s *ListPatrolReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
