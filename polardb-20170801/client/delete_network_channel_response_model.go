// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkChannelResponseBody) *DeleteNetworkChannelResponse
	GetBody() *DeleteNetworkChannelResponseBody
}

type DeleteNetworkChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkChannelResponse) GetBody() *DeleteNetworkChannelResponseBody {
	return s.Body
}

func (s *DeleteNetworkChannelResponse) SetHeaders(v map[string]*string) *DeleteNetworkChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkChannelResponse) SetStatusCode(v int32) *DeleteNetworkChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkChannelResponse) SetBody(v *DeleteNetworkChannelResponseBody) *DeleteNetworkChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
