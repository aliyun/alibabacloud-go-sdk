// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStartApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStartApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *BatchStartApplicationsResponseBody) *BatchStartApplicationsResponse
	GetBody() *BatchStartApplicationsResponseBody
}

type BatchStartApplicationsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStartApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStartApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStartApplicationsResponse) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStartApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStartApplicationsResponse) GetBody() *BatchStartApplicationsResponseBody {
	return s.Body
}

func (s *BatchStartApplicationsResponse) SetHeaders(v map[string]*string) *BatchStartApplicationsResponse {
	s.Headers = v
	return s
}

func (s *BatchStartApplicationsResponse) SetStatusCode(v int32) *BatchStartApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartApplicationsResponse) SetBody(v *BatchStartApplicationsResponseBody) *BatchStartApplicationsResponse {
	s.Body = v
	return s
}

func (s *BatchStartApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
