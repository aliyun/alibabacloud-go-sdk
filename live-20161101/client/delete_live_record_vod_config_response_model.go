// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordVodConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveRecordVodConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveRecordVodConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveRecordVodConfigResponseBody) *DeleteLiveRecordVodConfigResponse
	GetBody() *DeleteLiveRecordVodConfigResponseBody
}

type DeleteLiveRecordVodConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveRecordVodConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveRecordVodConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordVodConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordVodConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveRecordVodConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveRecordVodConfigResponse) GetBody() *DeleteLiveRecordVodConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveRecordVodConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveRecordVodConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRecordVodConfigResponse) SetStatusCode(v int32) *DeleteLiveRecordVodConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRecordVodConfigResponse) SetBody(v *DeleteLiveRecordVodConfigResponseBody) *DeleteLiveRecordVodConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveRecordVodConfigResponse) Validate() error {
	return dara.Validate(s)
}
