// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveSnapshotNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveSnapshotNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveSnapshotNotifyConfigResponseBody) *DeleteLiveSnapshotNotifyConfigResponse
	GetBody() *DeleteLiveSnapshotNotifyConfigResponseBody
}

type DeleteLiveSnapshotNotifyConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveSnapshotNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveSnapshotNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) GetBody() *DeleteLiveSnapshotNotifyConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveSnapshotNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) SetStatusCode(v int32) *DeleteLiveSnapshotNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) SetBody(v *DeleteLiveSnapshotNotifyConfigResponseBody) *DeleteLiveSnapshotNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveSnapshotNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
