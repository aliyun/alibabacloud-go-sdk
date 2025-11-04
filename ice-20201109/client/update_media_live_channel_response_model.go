// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaLiveChannelResponseBody) *UpdateMediaLiveChannelResponse
	GetBody() *UpdateMediaLiveChannelResponseBody
}

type UpdateMediaLiveChannelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaLiveChannelResponse) GetBody() *UpdateMediaLiveChannelResponseBody {
	return s.Body
}

func (s *UpdateMediaLiveChannelResponse) SetHeaders(v map[string]*string) *UpdateMediaLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaLiveChannelResponse) SetStatusCode(v int32) *UpdateMediaLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaLiveChannelResponse) SetBody(v *UpdateMediaLiveChannelResponseBody) *UpdateMediaLiveChannelResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
