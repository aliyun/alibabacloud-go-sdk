// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamsNotifyUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamsNotifyUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamsNotifyUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamsNotifyUrlConfigResponseBody) *DeleteLiveStreamsNotifyUrlConfigResponse
	GetBody() *DeleteLiveStreamsNotifyUrlConfigResponseBody
}

type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamsNotifyUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) GetBody() *DeleteLiveStreamsNotifyUrlConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) SetStatusCode(v int32) *DeleteLiveStreamsNotifyUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) SetBody(v *DeleteLiveStreamsNotifyUrlConfigResponseBody) *DeleteLiveStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
