// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFCTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFCTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFCTriggerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFCTriggerResponseBody) *UpdateFCTriggerResponse
	GetBody() *UpdateFCTriggerResponseBody
}

type UpdateFCTriggerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFCTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFCTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFCTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateFCTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFCTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFCTriggerResponse) GetBody() *UpdateFCTriggerResponseBody {
	return s.Body
}

func (s *UpdateFCTriggerResponse) SetHeaders(v map[string]*string) *UpdateFCTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateFCTriggerResponse) SetStatusCode(v int32) *UpdateFCTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFCTriggerResponse) SetBody(v *UpdateFCTriggerResponseBody) *UpdateFCTriggerResponse {
	s.Body = v
	return s
}

func (s *UpdateFCTriggerResponse) Validate() error {
	return dara.Validate(s)
}
