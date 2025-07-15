// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotDetectPornConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveSnapshotDetectPornConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveSnapshotDetectPornConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveSnapshotDetectPornConfigResponseBody) *DeleteLiveSnapshotDetectPornConfigResponse
	GetBody() *DeleteLiveSnapshotDetectPornConfigResponseBody
}

type DeleteLiveSnapshotDetectPornConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveSnapshotDetectPornConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveSnapshotDetectPornConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotDetectPornConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) GetBody() *DeleteLiveSnapshotDetectPornConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveSnapshotDetectPornConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) SetStatusCode(v int32) *DeleteLiveSnapshotDetectPornConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) SetBody(v *DeleteLiveSnapshotDetectPornConfigResponseBody) *DeleteLiveSnapshotDetectPornConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveSnapshotDetectPornConfigResponse) Validate() error {
	return dara.Validate(s)
}
