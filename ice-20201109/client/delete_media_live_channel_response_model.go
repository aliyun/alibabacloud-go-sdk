// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaLiveChannelResponseBody) *DeleteMediaLiveChannelResponse
	GetBody() *DeleteMediaLiveChannelResponseBody
}

type DeleteMediaLiveChannelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaLiveChannelResponse) GetBody() *DeleteMediaLiveChannelResponseBody {
	return s.Body
}

func (s *DeleteMediaLiveChannelResponse) SetHeaders(v map[string]*string) *DeleteMediaLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaLiveChannelResponse) SetStatusCode(v int32) *DeleteMediaLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaLiveChannelResponse) SetBody(v *DeleteMediaLiveChannelResponseBody) *DeleteMediaLiveChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
