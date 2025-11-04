// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMediaLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartMediaLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartMediaLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *StartMediaLiveChannelResponseBody) *StartMediaLiveChannelResponse
	GetBody() *StartMediaLiveChannelResponseBody
}

type StartMediaLiveChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMediaLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartMediaLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s StartMediaLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *StartMediaLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartMediaLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartMediaLiveChannelResponse) GetBody() *StartMediaLiveChannelResponseBody {
	return s.Body
}

func (s *StartMediaLiveChannelResponse) SetHeaders(v map[string]*string) *StartMediaLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *StartMediaLiveChannelResponse) SetStatusCode(v int32) *StartMediaLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMediaLiveChannelResponse) SetBody(v *StartMediaLiveChannelResponseBody) *StartMediaLiveChannelResponse {
	s.Body = v
	return s
}

func (s *StartMediaLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
