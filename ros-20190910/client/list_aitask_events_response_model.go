// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITaskEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAITaskEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAITaskEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListAITaskEventsResponseBody) *ListAITaskEventsResponse
	GetBody() *ListAITaskEventsResponseBody
}

type ListAITaskEventsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAITaskEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAITaskEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAITaskEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAITaskEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAITaskEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAITaskEventsResponse) GetBody() *ListAITaskEventsResponseBody {
	return s.Body
}

func (s *ListAITaskEventsResponse) SetHeaders(v map[string]*string) *ListAITaskEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAITaskEventsResponse) SetStatusCode(v int32) *ListAITaskEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAITaskEventsResponse) SetBody(v *ListAITaskEventsResponseBody) *ListAITaskEventsResponse {
	s.Body = v
	return s
}

func (s *ListAITaskEventsResponse) Validate() error {
	return dara.Validate(s)
}
