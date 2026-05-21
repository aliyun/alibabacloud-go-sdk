// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationEventsResponseBody) *ListOperationEventsResponse
	GetBody() *ListOperationEventsResponseBody
}

type ListOperationEventsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationEventsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationEventsResponse) GetBody() *ListOperationEventsResponseBody {
	return s.Body
}

func (s *ListOperationEventsResponse) SetHeaders(v map[string]*string) *ListOperationEventsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationEventsResponse) SetStatusCode(v int32) *ListOperationEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationEventsResponse) SetBody(v *ListOperationEventsResponseBody) *ListOperationEventsResponse {
	s.Body = v
	return s
}

func (s *ListOperationEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
