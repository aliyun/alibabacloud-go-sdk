// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetChannelResponseBody) *GetChannelResponse
	GetBody() *GetChannelResponseBody
}

type GetChannelResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChannelResponse) GoString() string {
	return s.String()
}

func (s *GetChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChannelResponse) GetBody() *GetChannelResponseBody {
	return s.Body
}

func (s *GetChannelResponse) SetHeaders(v map[string]*string) *GetChannelResponse {
	s.Headers = v
	return s
}

func (s *GetChannelResponse) SetStatusCode(v int32) *GetChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChannelResponse) SetBody(v *GetChannelResponseBody) *GetChannelResponse {
	s.Body = v
	return s
}

func (s *GetChannelResponse) Validate() error {
	return dara.Validate(s)
}
