// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventResponse
	GetStatusCode() *int32
	SetBody(v *GetEventResponseBody) *GetEventResponse
	GetBody() *GetEventResponseBody
}

type GetEventResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponse) GoString() string {
	return s.String()
}

func (s *GetEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventResponse) GetBody() *GetEventResponseBody {
	return s.Body
}

func (s *GetEventResponse) SetHeaders(v map[string]*string) *GetEventResponse {
	s.Headers = v
	return s
}

func (s *GetEventResponse) SetStatusCode(v int32) *GetEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventResponse) SetBody(v *GetEventResponseBody) *GetEventResponse {
	s.Body = v
	return s
}

func (s *GetEventResponse) Validate() error {
	return dara.Validate(s)
}
