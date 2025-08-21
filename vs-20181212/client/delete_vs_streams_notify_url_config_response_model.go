// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVsStreamsNotifyUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVsStreamsNotifyUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVsStreamsNotifyUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVsStreamsNotifyUrlConfigResponseBody) *DeleteVsStreamsNotifyUrlConfigResponse
	GetBody() *DeleteVsStreamsNotifyUrlConfigResponseBody
}

type DeleteVsStreamsNotifyUrlConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVsStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVsStreamsNotifyUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVsStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) GetBody() *DeleteVsStreamsNotifyUrlConfigResponseBody {
	return s.Body
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *DeleteVsStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) SetStatusCode(v int32) *DeleteVsStreamsNotifyUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) SetBody(v *DeleteVsStreamsNotifyUrlConfigResponseBody) *DeleteVsStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigResponse) Validate() error {
	return dara.Validate(s)
}
