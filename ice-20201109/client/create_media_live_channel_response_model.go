// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaLiveChannelResponseBody) *CreateMediaLiveChannelResponse
	GetBody() *CreateMediaLiveChannelResponseBody
}

type CreateMediaLiveChannelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaLiveChannelResponse) GetBody() *CreateMediaLiveChannelResponseBody {
	return s.Body
}

func (s *CreateMediaLiveChannelResponse) SetHeaders(v map[string]*string) *CreateMediaLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaLiveChannelResponse) SetStatusCode(v int32) *CreateMediaLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaLiveChannelResponse) SetBody(v *CreateMediaLiveChannelResponseBody) *CreateMediaLiveChannelResponse {
	s.Body = v
	return s
}

func (s *CreateMediaLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
