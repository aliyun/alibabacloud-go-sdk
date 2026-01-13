// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLaboratoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLaboratoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLaboratoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListLaboratoriesResponseBody) *ListLaboratoriesResponse
	GetBody() *ListLaboratoriesResponseBody
}

type ListLaboratoriesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLaboratoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLaboratoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLaboratoriesResponse) GoString() string {
	return s.String()
}

func (s *ListLaboratoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLaboratoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLaboratoriesResponse) GetBody() *ListLaboratoriesResponseBody {
	return s.Body
}

func (s *ListLaboratoriesResponse) SetHeaders(v map[string]*string) *ListLaboratoriesResponse {
	s.Headers = v
	return s
}

func (s *ListLaboratoriesResponse) SetStatusCode(v int32) *ListLaboratoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLaboratoriesResponse) SetBody(v *ListLaboratoriesResponseBody) *ListLaboratoriesResponse {
	s.Body = v
	return s
}

func (s *ListLaboratoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
