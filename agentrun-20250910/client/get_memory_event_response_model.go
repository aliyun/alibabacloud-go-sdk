// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryEventResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryEventResponseBody) *GetMemoryEventResponse
	GetBody() *GetMemoryEventResponseBody
}

type GetMemoryEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryEventResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryEventResponse) GetBody() *GetMemoryEventResponseBody {
	return s.Body
}

func (s *GetMemoryEventResponse) SetHeaders(v map[string]*string) *GetMemoryEventResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryEventResponse) SetStatusCode(v int32) *GetMemoryEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryEventResponse) SetBody(v *GetMemoryEventResponseBody) *GetMemoryEventResponse {
	s.Body = v
	return s
}

func (s *GetMemoryEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
