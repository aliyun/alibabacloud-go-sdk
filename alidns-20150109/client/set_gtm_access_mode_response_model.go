// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetGtmAccessModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetGtmAccessModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetGtmAccessModeResponse
	GetStatusCode() *int32
	SetBody(v *SetGtmAccessModeResponseBody) *SetGtmAccessModeResponse
	GetBody() *SetGtmAccessModeResponseBody
}

type SetGtmAccessModeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetGtmAccessModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetGtmAccessModeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetGtmAccessModeResponse) GoString() string {
	return s.String()
}

func (s *SetGtmAccessModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetGtmAccessModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetGtmAccessModeResponse) GetBody() *SetGtmAccessModeResponseBody {
	return s.Body
}

func (s *SetGtmAccessModeResponse) SetHeaders(v map[string]*string) *SetGtmAccessModeResponse {
	s.Headers = v
	return s
}

func (s *SetGtmAccessModeResponse) SetStatusCode(v int32) *SetGtmAccessModeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetGtmAccessModeResponse) SetBody(v *SetGtmAccessModeResponseBody) *SetGtmAccessModeResponse {
	s.Body = v
	return s
}

func (s *SetGtmAccessModeResponse) Validate() error {
	return dara.Validate(s)
}
