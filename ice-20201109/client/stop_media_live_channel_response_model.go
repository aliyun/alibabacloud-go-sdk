// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMediaLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopMediaLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopMediaLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *StopMediaLiveChannelResponseBody) *StopMediaLiveChannelResponse
	GetBody() *StopMediaLiveChannelResponseBody
}

type StopMediaLiveChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMediaLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopMediaLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s StopMediaLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *StopMediaLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopMediaLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopMediaLiveChannelResponse) GetBody() *StopMediaLiveChannelResponseBody {
	return s.Body
}

func (s *StopMediaLiveChannelResponse) SetHeaders(v map[string]*string) *StopMediaLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *StopMediaLiveChannelResponse) SetStatusCode(v int32) *StopMediaLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMediaLiveChannelResponse) SetBody(v *StopMediaLiveChannelResponseBody) *StopMediaLiveChannelResponse {
	s.Body = v
	return s
}

func (s *StopMediaLiveChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
