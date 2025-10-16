// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMemoryEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMemoryEventResponse
	GetStatusCode() *int32
	SetBody(v *ListMemoryEventResponseBody) *ListMemoryEventResponse
	GetBody() *ListMemoryEventResponseBody
}

type ListMemoryEventResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryEventResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMemoryEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMemoryEventResponse) GetBody() *ListMemoryEventResponseBody {
	return s.Body
}

func (s *ListMemoryEventResponse) SetHeaders(v map[string]*string) *ListMemoryEventResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryEventResponse) SetStatusCode(v int32) *ListMemoryEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryEventResponse) SetBody(v *ListMemoryEventResponseBody) *ListMemoryEventResponse {
	s.Body = v
	return s
}

func (s *ListMemoryEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
