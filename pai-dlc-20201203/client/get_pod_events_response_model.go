// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPodEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPodEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPodEventsResponse
	GetStatusCode() *int32
	SetBody(v *GetPodEventsResponseBody) *GetPodEventsResponse
	GetBody() *GetPodEventsResponseBody
}

type GetPodEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPodEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPodEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPodEventsResponse) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPodEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPodEventsResponse) GetBody() *GetPodEventsResponseBody {
	return s.Body
}

func (s *GetPodEventsResponse) SetHeaders(v map[string]*string) *GetPodEventsResponse {
	s.Headers = v
	return s
}

func (s *GetPodEventsResponse) SetStatusCode(v int32) *GetPodEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPodEventsResponse) SetBody(v *GetPodEventsResponseBody) *GetPodEventsResponse {
	s.Body = v
	return s
}

func (s *GetPodEventsResponse) Validate() error {
	return dara.Validate(s)
}
