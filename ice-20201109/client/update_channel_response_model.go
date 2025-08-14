// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateChannelResponseBody) *UpdateChannelResponse
	GetBody() *UpdateChannelResponseBody
}

type UpdateChannelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateChannelResponse) GetBody() *UpdateChannelResponseBody {
	return s.Body
}

func (s *UpdateChannelResponse) SetHeaders(v map[string]*string) *UpdateChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateChannelResponse) SetStatusCode(v int32) *UpdateChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChannelResponse) SetBody(v *UpdateChannelResponseBody) *UpdateChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateChannelResponse) Validate() error {
	return dara.Validate(s)
}
