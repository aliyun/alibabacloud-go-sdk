// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessAssignmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAccessAssignmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAccessAssignmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListAccessAssignmentsResponseBody) *ListAccessAssignmentsResponse
	GetBody() *ListAccessAssignmentsResponseBody
}

type ListAccessAssignmentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessAssignmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessAssignmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAccessAssignmentsResponse) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAccessAssignmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAccessAssignmentsResponse) GetBody() *ListAccessAssignmentsResponseBody {
	return s.Body
}

func (s *ListAccessAssignmentsResponse) SetHeaders(v map[string]*string) *ListAccessAssignmentsResponse {
	s.Headers = v
	return s
}

func (s *ListAccessAssignmentsResponse) SetStatusCode(v int32) *ListAccessAssignmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessAssignmentsResponse) SetBody(v *ListAccessAssignmentsResponseBody) *ListAccessAssignmentsResponse {
	s.Body = v
	return s
}

func (s *ListAccessAssignmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
