// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAppSnapshotConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAppSnapshotConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAppSnapshotConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAppSnapshotConfigResponseBody) *DeleteLiveAppSnapshotConfigResponse
	GetBody() *DeleteLiveAppSnapshotConfigResponseBody
}

type DeleteLiveAppSnapshotConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAppSnapshotConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAppSnapshotConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAppSnapshotConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppSnapshotConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAppSnapshotConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAppSnapshotConfigResponse) GetBody() *DeleteLiveAppSnapshotConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveAppSnapshotConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveAppSnapshotConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAppSnapshotConfigResponse) SetStatusCode(v int32) *DeleteLiveAppSnapshotConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAppSnapshotConfigResponse) SetBody(v *DeleteLiveAppSnapshotConfigResponseBody) *DeleteLiveAppSnapshotConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAppSnapshotConfigResponse) Validate() error {
	return dara.Validate(s)
}
