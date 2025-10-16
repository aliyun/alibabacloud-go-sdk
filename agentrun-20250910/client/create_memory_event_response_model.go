// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemoryEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemoryEventResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemoryEventResponseBody) *CreateMemoryEventResponse
	GetBody() *CreateMemoryEventResponseBody
}

type CreateMemoryEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryEventResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryEventResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemoryEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemoryEventResponse) GetBody() *CreateMemoryEventResponseBody {
	return s.Body
}

func (s *CreateMemoryEventResponse) SetHeaders(v map[string]*string) *CreateMemoryEventResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryEventResponse) SetStatusCode(v int32) *CreateMemoryEventResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryEventResponse) SetBody(v *CreateMemoryEventResponseBody) *CreateMemoryEventResponse {
	s.Body = v
	return s
}

func (s *CreateMemoryEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
