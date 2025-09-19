// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableHoneypotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableHoneypotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableHoneypotResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableHoneypotResponseBody) *ListAvailableHoneypotResponse
	GetBody() *ListAvailableHoneypotResponseBody
}

type ListAvailableHoneypotResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableHoneypotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableHoneypotResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableHoneypotResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableHoneypotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableHoneypotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableHoneypotResponse) GetBody() *ListAvailableHoneypotResponseBody {
	return s.Body
}

func (s *ListAvailableHoneypotResponse) SetHeaders(v map[string]*string) *ListAvailableHoneypotResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableHoneypotResponse) SetStatusCode(v int32) *ListAvailableHoneypotResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableHoneypotResponse) SetBody(v *ListAvailableHoneypotResponseBody) *ListAvailableHoneypotResponse {
	s.Body = v
	return s
}

func (s *ListAvailableHoneypotResponse) Validate() error {
	return dara.Validate(s)
}
