// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventSourceResponseBody) *CreateEventSourceResponse
	GetBody() *CreateEventSourceResponseBody
}

type CreateEventSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateEventSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventSourceResponse) GetBody() *CreateEventSourceResponseBody {
	return s.Body
}

func (s *CreateEventSourceResponse) SetHeaders(v map[string]*string) *CreateEventSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateEventSourceResponse) SetStatusCode(v int32) *CreateEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventSourceResponse) SetBody(v *CreateEventSourceResponseBody) *CreateEventSourceResponse {
	s.Body = v
	return s
}

func (s *CreateEventSourceResponse) Validate() error {
	return dara.Validate(s)
}
