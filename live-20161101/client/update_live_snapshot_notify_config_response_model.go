// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveSnapshotNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveSnapshotNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveSnapshotNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveSnapshotNotifyConfigResponseBody) *UpdateLiveSnapshotNotifyConfigResponse
	GetBody() *UpdateLiveSnapshotNotifyConfigResponseBody
}

type UpdateLiveSnapshotNotifyConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveSnapshotNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveSnapshotNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveSnapshotNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) GetBody() *UpdateLiveSnapshotNotifyConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveSnapshotNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) SetStatusCode(v int32) *UpdateLiveSnapshotNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) SetBody(v *UpdateLiveSnapshotNotifyConfigResponseBody) *UpdateLiveSnapshotNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveSnapshotNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
