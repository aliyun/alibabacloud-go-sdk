// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaLiveChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaLiveChannelResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaLiveChannelResponseBody) *GetMediaLiveChannelResponse
	GetBody() *GetMediaLiveChannelResponseBody
}

type GetMediaLiveChannelResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaLiveChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *GetMediaLiveChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaLiveChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaLiveChannelResponse) GetBody() *GetMediaLiveChannelResponseBody {
	return s.Body
}

func (s *GetMediaLiveChannelResponse) SetHeaders(v map[string]*string) *GetMediaLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *GetMediaLiveChannelResponse) SetStatusCode(v int32) *GetMediaLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaLiveChannelResponse) SetBody(v *GetMediaLiveChannelResponseBody) *GetMediaLiveChannelResponse {
	s.Body = v
	return s
}

func (s *GetMediaLiveChannelResponse) Validate() error {
	return dara.Validate(s)
}
