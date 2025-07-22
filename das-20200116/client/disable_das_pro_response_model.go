// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDasProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDasProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDasProResponse
	GetStatusCode() *int32
	SetBody(v *DisableDasProResponseBody) *DisableDasProResponse
	GetBody() *DisableDasProResponseBody
}

type DisableDasProResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDasProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDasProResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDasProResponse) GoString() string {
	return s.String()
}

func (s *DisableDasProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDasProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDasProResponse) GetBody() *DisableDasProResponseBody {
	return s.Body
}

func (s *DisableDasProResponse) SetHeaders(v map[string]*string) *DisableDasProResponse {
	s.Headers = v
	return s
}

func (s *DisableDasProResponse) SetStatusCode(v int32) *DisableDasProResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDasProResponse) SetBody(v *DisableDasProResponseBody) *DisableDasProResponse {
	s.Body = v
	return s
}

func (s *DisableDasProResponse) Validate() error {
	return dara.Validate(s)
}
