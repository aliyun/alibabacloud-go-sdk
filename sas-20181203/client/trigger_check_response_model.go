// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerCheckResponse
	GetStatusCode() *int32
	SetBody(v *TriggerCheckResponseBody) *TriggerCheckResponse
	GetBody() *TriggerCheckResponseBody
}

type TriggerCheckResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerCheckResponse) GoString() string {
	return s.String()
}

func (s *TriggerCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerCheckResponse) GetBody() *TriggerCheckResponseBody {
	return s.Body
}

func (s *TriggerCheckResponse) SetHeaders(v map[string]*string) *TriggerCheckResponse {
	s.Headers = v
	return s
}

func (s *TriggerCheckResponse) SetStatusCode(v int32) *TriggerCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerCheckResponse) SetBody(v *TriggerCheckResponseBody) *TriggerCheckResponse {
	s.Body = v
	return s
}

func (s *TriggerCheckResponse) Validate() error {
	return dara.Validate(s)
}
