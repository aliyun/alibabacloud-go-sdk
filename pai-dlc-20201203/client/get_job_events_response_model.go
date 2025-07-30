// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJobEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJobEventsResponse
	GetStatusCode() *int32
	SetBody(v *GetJobEventsResponseBody) *GetJobEventsResponse
	GetBody() *GetJobEventsResponseBody
}

type GetJobEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJobEventsResponse) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJobEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJobEventsResponse) GetBody() *GetJobEventsResponseBody {
	return s.Body
}

func (s *GetJobEventsResponse) SetHeaders(v map[string]*string) *GetJobEventsResponse {
	s.Headers = v
	return s
}

func (s *GetJobEventsResponse) SetStatusCode(v int32) *GetJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobEventsResponse) SetBody(v *GetJobEventsResponseBody) *GetJobEventsResponse {
	s.Body = v
	return s
}

func (s *GetJobEventsResponse) Validate() error {
	return dara.Validate(s)
}
