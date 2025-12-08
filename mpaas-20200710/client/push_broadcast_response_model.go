// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBroadcastResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushBroadcastResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushBroadcastResponse
	GetStatusCode() *int32
	SetBody(v *PushBroadcastResponseBody) *PushBroadcastResponse
	GetBody() *PushBroadcastResponseBody
}

type PushBroadcastResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushBroadcastResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushBroadcastResponse) String() string {
	return dara.Prettify(s)
}

func (s PushBroadcastResponse) GoString() string {
	return s.String()
}

func (s *PushBroadcastResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushBroadcastResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushBroadcastResponse) GetBody() *PushBroadcastResponseBody {
	return s.Body
}

func (s *PushBroadcastResponse) SetHeaders(v map[string]*string) *PushBroadcastResponse {
	s.Headers = v
	return s
}

func (s *PushBroadcastResponse) SetStatusCode(v int32) *PushBroadcastResponse {
	s.StatusCode = &v
	return s
}

func (s *PushBroadcastResponse) SetBody(v *PushBroadcastResponseBody) *PushBroadcastResponse {
	s.Body = v
	return s
}

func (s *PushBroadcastResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
