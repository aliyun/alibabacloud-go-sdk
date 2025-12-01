// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendCustomEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendCustomEventResponse
	GetStatusCode() *int32
	SetBody(v *SendCustomEventResponseBody) *SendCustomEventResponse
	GetBody() *SendCustomEventResponseBody
}

type SendCustomEventResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendCustomEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendCustomEventResponse) String() string {
	return dara.Prettify(s)
}

func (s SendCustomEventResponse) GoString() string {
	return s.String()
}

func (s *SendCustomEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendCustomEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendCustomEventResponse) GetBody() *SendCustomEventResponseBody {
	return s.Body
}

func (s *SendCustomEventResponse) SetHeaders(v map[string]*string) *SendCustomEventResponse {
	s.Headers = v
	return s
}

func (s *SendCustomEventResponse) SetStatusCode(v int32) *SendCustomEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCustomEventResponse) SetBody(v *SendCustomEventResponseBody) *SendCustomEventResponse {
	s.Body = v
	return s
}

func (s *SendCustomEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
