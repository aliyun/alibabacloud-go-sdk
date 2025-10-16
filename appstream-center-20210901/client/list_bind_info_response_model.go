// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBindInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBindInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListBindInfoResponseBody) *ListBindInfoResponse
	GetBody() *ListBindInfoResponseBody
}

type ListBindInfoResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBindInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBindInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBindInfoResponse) GoString() string {
	return s.String()
}

func (s *ListBindInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBindInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBindInfoResponse) GetBody() *ListBindInfoResponseBody {
	return s.Body
}

func (s *ListBindInfoResponse) SetHeaders(v map[string]*string) *ListBindInfoResponse {
	s.Headers = v
	return s
}

func (s *ListBindInfoResponse) SetStatusCode(v int32) *ListBindInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBindInfoResponse) SetBody(v *ListBindInfoResponseBody) *ListBindInfoResponse {
	s.Body = v
	return s
}

func (s *ListBindInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
