// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventSubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventSubscribeResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventSubscribeResponseBody) *CreateEventSubscribeResponse
	GetBody() *CreateEventSubscribeResponseBody
}

type CreateEventSubscribeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventSubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSubscribeResponse) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventSubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventSubscribeResponse) GetBody() *CreateEventSubscribeResponseBody {
	return s.Body
}

func (s *CreateEventSubscribeResponse) SetHeaders(v map[string]*string) *CreateEventSubscribeResponse {
	s.Headers = v
	return s
}

func (s *CreateEventSubscribeResponse) SetStatusCode(v int32) *CreateEventSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventSubscribeResponse) SetBody(v *CreateEventSubscribeResponseBody) *CreateEventSubscribeResponse {
	s.Body = v
	return s
}

func (s *CreateEventSubscribeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
