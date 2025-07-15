// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBargeInCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BargeInCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BargeInCallResponse
	GetStatusCode() *int32
	SetBody(v *BargeInCallResponseBody) *BargeInCallResponse
	GetBody() *BargeInCallResponseBody
}

type BargeInCallResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BargeInCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BargeInCallResponse) String() string {
	return dara.Prettify(s)
}

func (s BargeInCallResponse) GoString() string {
	return s.String()
}

func (s *BargeInCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BargeInCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BargeInCallResponse) GetBody() *BargeInCallResponseBody {
	return s.Body
}

func (s *BargeInCallResponse) SetHeaders(v map[string]*string) *BargeInCallResponse {
	s.Headers = v
	return s
}

func (s *BargeInCallResponse) SetStatusCode(v int32) *BargeInCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BargeInCallResponse) SetBody(v *BargeInCallResponseBody) *BargeInCallResponse {
	s.Body = v
	return s
}

func (s *BargeInCallResponse) Validate() error {
	return dara.Validate(s)
}
