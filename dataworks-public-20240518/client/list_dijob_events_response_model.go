// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDIJobEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDIJobEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListDIJobEventsResponseBody) *ListDIJobEventsResponse
	GetBody() *ListDIJobEventsResponseBody
}

type ListDIJobEventsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobEventsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDIJobEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDIJobEventsResponse) GetBody() *ListDIJobEventsResponseBody {
	return s.Body
}

func (s *ListDIJobEventsResponse) SetHeaders(v map[string]*string) *ListDIJobEventsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobEventsResponse) SetStatusCode(v int32) *ListDIJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobEventsResponse) SetBody(v *ListDIJobEventsResponseBody) *ListDIJobEventsResponse {
	s.Body = v
	return s
}

func (s *ListDIJobEventsResponse) Validate() error {
	return dara.Validate(s)
}
