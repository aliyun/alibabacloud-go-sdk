// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutCustomEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutCustomEventResponse
	GetStatusCode() *int32
	SetBody(v *PutCustomEventResponseBody) *PutCustomEventResponse
	GetBody() *PutCustomEventResponseBody
}

type PutCustomEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutCustomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutCustomEventResponse) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventResponse) GoString() string {
	return s.String()
}

func (s *PutCustomEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutCustomEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutCustomEventResponse) GetBody() *PutCustomEventResponseBody {
	return s.Body
}

func (s *PutCustomEventResponse) SetHeaders(v map[string]*string) *PutCustomEventResponse {
	s.Headers = v
	return s
}

func (s *PutCustomEventResponse) SetStatusCode(v int32) *PutCustomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *PutCustomEventResponse) SetBody(v *PutCustomEventResponseBody) *PutCustomEventResponse {
	s.Body = v
	return s
}

func (s *PutCustomEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
