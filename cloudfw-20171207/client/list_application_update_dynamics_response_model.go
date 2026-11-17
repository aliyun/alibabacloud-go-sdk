// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationUpdateDynamicsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationUpdateDynamicsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationUpdateDynamicsResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationUpdateDynamicsResponseBody) *ListApplicationUpdateDynamicsResponse
	GetBody() *ListApplicationUpdateDynamicsResponseBody
}

type ListApplicationUpdateDynamicsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationUpdateDynamicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationUpdateDynamicsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationUpdateDynamicsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationUpdateDynamicsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationUpdateDynamicsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationUpdateDynamicsResponse) GetBody() *ListApplicationUpdateDynamicsResponseBody {
	return s.Body
}

func (s *ListApplicationUpdateDynamicsResponse) SetHeaders(v map[string]*string) *ListApplicationUpdateDynamicsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationUpdateDynamicsResponse) SetStatusCode(v int32) *ListApplicationUpdateDynamicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationUpdateDynamicsResponse) SetBody(v *ListApplicationUpdateDynamicsResponseBody) *ListApplicationUpdateDynamicsResponse {
	s.Body = v
	return s
}

func (s *ListApplicationUpdateDynamicsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
