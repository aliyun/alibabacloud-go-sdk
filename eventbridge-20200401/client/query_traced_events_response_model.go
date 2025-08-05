// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTracedEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTracedEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTracedEventsResponse
	GetStatusCode() *int32
	SetBody(v *QueryTracedEventsResponseBody) *QueryTracedEventsResponse
	GetBody() *QueryTracedEventsResponseBody
}

type QueryTracedEventsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryTracedEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTracedEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTracedEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryTracedEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTracedEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTracedEventsResponse) GetBody() *QueryTracedEventsResponseBody {
	return s.Body
}

func (s *QueryTracedEventsResponse) SetHeaders(v map[string]*string) *QueryTracedEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryTracedEventsResponse) SetStatusCode(v int32) *QueryTracedEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTracedEventsResponse) SetBody(v *QueryTracedEventsResponseBody) *QueryTracedEventsResponse {
	s.Body = v
	return s
}

func (s *QueryTracedEventsResponse) Validate() error {
	return dara.Validate(s)
}
