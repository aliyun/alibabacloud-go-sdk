// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamsNotifyUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveStreamsNotifyUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveStreamsNotifyUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveStreamsNotifyUrlConfigResponseBody) *SetLiveStreamsNotifyUrlConfigResponse
	GetBody() *SetLiveStreamsNotifyUrlConfigResponseBody
}

type SetLiveStreamsNotifyUrlConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveStreamsNotifyUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) GetBody() *SetLiveStreamsNotifyUrlConfigResponseBody {
	return s.Body
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *SetLiveStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) SetStatusCode(v int32) *SetLiveStreamsNotifyUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) SetBody(v *SetLiveStreamsNotifyUrlConfigResponseBody) *SetLiveStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigResponse) Validate() error {
	return dara.Validate(s)
}
