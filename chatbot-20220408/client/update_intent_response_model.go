// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIntentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIntentResponseBody) *UpdateIntentResponse
	GetBody() *UpdateIntentResponseBody
}

type UpdateIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntentResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIntentResponse) GetBody() *UpdateIntentResponseBody {
	return s.Body
}

func (s *UpdateIntentResponse) SetHeaders(v map[string]*string) *UpdateIntentResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntentResponse) SetStatusCode(v int32) *UpdateIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntentResponse) SetBody(v *UpdateIntentResponseBody) *UpdateIntentResponse {
	s.Body = v
	return s
}

func (s *UpdateIntentResponse) Validate() error {
	return dara.Validate(s)
}
