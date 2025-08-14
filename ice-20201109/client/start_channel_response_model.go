// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartChannelResponse
	GetStatusCode() *int32
	SetBody(v *StartChannelResponseBody) *StartChannelResponse
	GetBody() *StartChannelResponseBody
}

type StartChannelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s StartChannelResponse) GoString() string {
	return s.String()
}

func (s *StartChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartChannelResponse) GetBody() *StartChannelResponseBody {
	return s.Body
}

func (s *StartChannelResponse) SetHeaders(v map[string]*string) *StartChannelResponse {
	s.Headers = v
	return s
}

func (s *StartChannelResponse) SetStatusCode(v int32) *StartChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *StartChannelResponse) SetBody(v *StartChannelResponseBody) *StartChannelResponse {
	s.Body = v
	return s
}

func (s *StartChannelResponse) Validate() error {
	return dara.Validate(s)
}
