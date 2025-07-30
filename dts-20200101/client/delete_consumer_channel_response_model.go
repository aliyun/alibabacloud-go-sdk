// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConsumerChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConsumerChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConsumerChannelResponseBody) *DeleteConsumerChannelResponse
	GetBody() *DeleteConsumerChannelResponseBody
}

type DeleteConsumerChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConsumerChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConsumerChannelResponse) GetBody() *DeleteConsumerChannelResponseBody {
	return s.Body
}

func (s *DeleteConsumerChannelResponse) SetHeaders(v map[string]*string) *DeleteConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerChannelResponse) SetStatusCode(v int32) *DeleteConsumerChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerChannelResponse) SetBody(v *DeleteConsumerChannelResponseBody) *DeleteConsumerChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteConsumerChannelResponse) Validate() error {
	return dara.Validate(s)
}
