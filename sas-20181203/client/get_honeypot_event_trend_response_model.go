// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotEventTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoneypotEventTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoneypotEventTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetHoneypotEventTrendResponseBody) *GetHoneypotEventTrendResponse
	GetBody() *GetHoneypotEventTrendResponseBody
}

type GetHoneypotEventTrendResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoneypotEventTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoneypotEventTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotEventTrendResponse) GoString() string {
	return s.String()
}

func (s *GetHoneypotEventTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoneypotEventTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoneypotEventTrendResponse) GetBody() *GetHoneypotEventTrendResponseBody {
	return s.Body
}

func (s *GetHoneypotEventTrendResponse) SetHeaders(v map[string]*string) *GetHoneypotEventTrendResponse {
	s.Headers = v
	return s
}

func (s *GetHoneypotEventTrendResponse) SetStatusCode(v int32) *GetHoneypotEventTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoneypotEventTrendResponse) SetBody(v *GetHoneypotEventTrendResponseBody) *GetHoneypotEventTrendResponse {
	s.Body = v
	return s
}

func (s *GetHoneypotEventTrendResponse) Validate() error {
	return dara.Validate(s)
}
