// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnsubscribeEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnsubscribeEventResponse
	GetStatusCode() *int32
	SetBody(v *UnsubscribeEventResponseBody) *UnsubscribeEventResponse
	GetBody() *UnsubscribeEventResponseBody
}

type UnsubscribeEventResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeEventResponse) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnsubscribeEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnsubscribeEventResponse) GetBody() *UnsubscribeEventResponseBody {
	return s.Body
}

func (s *UnsubscribeEventResponse) SetHeaders(v map[string]*string) *UnsubscribeEventResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeEventResponse) SetStatusCode(v int32) *UnsubscribeEventResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeEventResponse) SetBody(v *UnsubscribeEventResponseBody) *UnsubscribeEventResponse {
	s.Body = v
	return s
}

func (s *UnsubscribeEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
