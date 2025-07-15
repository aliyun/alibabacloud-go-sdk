// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTakeBreakResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TakeBreakResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TakeBreakResponse
	GetStatusCode() *int32
	SetBody(v *TakeBreakResponseBody) *TakeBreakResponse
	GetBody() *TakeBreakResponseBody
}

type TakeBreakResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TakeBreakResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TakeBreakResponse) String() string {
	return dara.Prettify(s)
}

func (s TakeBreakResponse) GoString() string {
	return s.String()
}

func (s *TakeBreakResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TakeBreakResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TakeBreakResponse) GetBody() *TakeBreakResponseBody {
	return s.Body
}

func (s *TakeBreakResponse) SetHeaders(v map[string]*string) *TakeBreakResponse {
	s.Headers = v
	return s
}

func (s *TakeBreakResponse) SetStatusCode(v int32) *TakeBreakResponse {
	s.StatusCode = &v
	return s
}

func (s *TakeBreakResponse) SetBody(v *TakeBreakResponseBody) *TakeBreakResponse {
	s.Body = v
	return s
}

func (s *TakeBreakResponse) Validate() error {
	return dara.Validate(s)
}
