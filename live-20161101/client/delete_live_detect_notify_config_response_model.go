// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDetectNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveDetectNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveDetectNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveDetectNotifyConfigResponseBody) *DeleteLiveDetectNotifyConfigResponse
	GetBody() *DeleteLiveDetectNotifyConfigResponseBody
}

type DeleteLiveDetectNotifyConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveDetectNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveDetectNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveDetectNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveDetectNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveDetectNotifyConfigResponse) GetBody() *DeleteLiveDetectNotifyConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveDetectNotifyConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveDetectNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveDetectNotifyConfigResponse) SetStatusCode(v int32) *DeleteLiveDetectNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveDetectNotifyConfigResponse) SetBody(v *DeleteLiveDetectNotifyConfigResponseBody) *DeleteLiveDetectNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveDetectNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
