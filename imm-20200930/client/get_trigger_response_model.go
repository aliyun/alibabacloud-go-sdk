// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTriggerResponse
	GetStatusCode() *int32
	SetBody(v *GetTriggerResponseBody) *GetTriggerResponse
	GetBody() *GetTriggerResponseBody
}

type GetTriggerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTriggerResponse) GoString() string {
	return s.String()
}

func (s *GetTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTriggerResponse) GetBody() *GetTriggerResponseBody {
	return s.Body
}

func (s *GetTriggerResponse) SetHeaders(v map[string]*string) *GetTriggerResponse {
	s.Headers = v
	return s
}

func (s *GetTriggerResponse) SetStatusCode(v int32) *GetTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTriggerResponse) SetBody(v *GetTriggerResponseBody) *GetTriggerResponse {
	s.Body = v
	return s
}

func (s *GetTriggerResponse) Validate() error {
	return dara.Validate(s)
}
