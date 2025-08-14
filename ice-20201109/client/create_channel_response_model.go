// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateChannelResponseBody) *CreateChannelResponse
	GetBody() *CreateChannelResponseBody
}

type CreateChannelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateChannelResponse) GetBody() *CreateChannelResponseBody {
	return s.Body
}

func (s *CreateChannelResponse) SetHeaders(v map[string]*string) *CreateChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateChannelResponse) SetStatusCode(v int32) *CreateChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChannelResponse) SetBody(v *CreateChannelResponseBody) *CreateChannelResponse {
	s.Body = v
	return s
}

func (s *CreateChannelResponse) Validate() error {
	return dara.Validate(s)
}
