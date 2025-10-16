// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersistentAppInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPersistentAppInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPersistentAppInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListPersistentAppInstancesResponseBody) *ListPersistentAppInstancesResponse
	GetBody() *ListPersistentAppInstancesResponseBody
}

type ListPersistentAppInstancesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPersistentAppInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPersistentAppInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPersistentAppInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListPersistentAppInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPersistentAppInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPersistentAppInstancesResponse) GetBody() *ListPersistentAppInstancesResponseBody {
	return s.Body
}

func (s *ListPersistentAppInstancesResponse) SetHeaders(v map[string]*string) *ListPersistentAppInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListPersistentAppInstancesResponse) SetStatusCode(v int32) *ListPersistentAppInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPersistentAppInstancesResponse) SetBody(v *ListPersistentAppInstancesResponseBody) *ListPersistentAppInstancesResponse {
	s.Body = v
	return s
}

func (s *ListPersistentAppInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
