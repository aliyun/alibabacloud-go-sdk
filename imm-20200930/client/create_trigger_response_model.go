// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTriggerResponse
	GetStatusCode() *int32
	SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse
	GetBody() *CreateTriggerResponseBody
}

type CreateTriggerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTriggerResponse) GetBody() *CreateTriggerResponseBody {
	return s.Body
}

func (s *CreateTriggerResponse) SetHeaders(v map[string]*string) *CreateTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateTriggerResponse) SetStatusCode(v int32) *CreateTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTriggerResponse) SetBody(v *CreateTriggerResponseBody) *CreateTriggerResponse {
	s.Body = v
	return s
}

func (s *CreateTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
