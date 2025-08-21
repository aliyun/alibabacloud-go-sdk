// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIntentResponse
	GetStatusCode() *int32
	SetBody(v *CreateIntentResponseBody) *CreateIntentResponse
	GetBody() *CreateIntentResponseBody
}

type CreateIntentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentResponse) GoString() string {
	return s.String()
}

func (s *CreateIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIntentResponse) GetBody() *CreateIntentResponseBody {
	return s.Body
}

func (s *CreateIntentResponse) SetHeaders(v map[string]*string) *CreateIntentResponse {
	s.Headers = v
	return s
}

func (s *CreateIntentResponse) SetStatusCode(v int32) *CreateIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntentResponse) SetBody(v *CreateIntentResponseBody) *CreateIntentResponse {
	s.Body = v
	return s
}

func (s *CreateIntentResponse) Validate() error {
	return dara.Validate(s)
}
