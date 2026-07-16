// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse
	GetBody() *ListEnvironmentsResponseBody
}

type ListEnvironmentsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentsResponse) GetBody() *ListEnvironmentsResponseBody {
	return s.Body
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetStatusCode(v int32) *ListEnvironmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
