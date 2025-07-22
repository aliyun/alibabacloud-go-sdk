// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceEventsResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceEventsResponseBody) *GetInstanceEventsResponse
	GetBody() *GetInstanceEventsResponseBody
}

type GetInstanceEventsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEventsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceEventsResponse) GetBody() *GetInstanceEventsResponseBody {
	return s.Body
}

func (s *GetInstanceEventsResponse) SetHeaders(v map[string]*string) *GetInstanceEventsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceEventsResponse) SetStatusCode(v int32) *GetInstanceEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceEventsResponse) SetBody(v *GetInstanceEventsResponseBody) *GetInstanceEventsResponse {
	s.Body = v
	return s
}

func (s *GetInstanceEventsResponse) Validate() error {
	return dara.Validate(s)
}
