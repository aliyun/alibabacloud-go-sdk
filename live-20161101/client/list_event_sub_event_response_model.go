// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventSubEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventSubEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventSubEventResponse
	GetStatusCode() *int32
	SetBody(v *ListEventSubEventResponseBody) *ListEventSubEventResponse
	GetBody() *ListEventSubEventResponseBody
}

type ListEventSubEventResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventSubEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventSubEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventSubEventResponse) GoString() string {
	return s.String()
}

func (s *ListEventSubEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventSubEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventSubEventResponse) GetBody() *ListEventSubEventResponseBody {
	return s.Body
}

func (s *ListEventSubEventResponse) SetHeaders(v map[string]*string) *ListEventSubEventResponse {
	s.Headers = v
	return s
}

func (s *ListEventSubEventResponse) SetStatusCode(v int32) *ListEventSubEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventSubEventResponse) SetBody(v *ListEventSubEventResponseBody) *ListEventSubEventResponse {
	s.Body = v
	return s
}

func (s *ListEventSubEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
