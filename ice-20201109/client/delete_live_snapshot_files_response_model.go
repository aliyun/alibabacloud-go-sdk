// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveSnapshotFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveSnapshotFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveSnapshotFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveSnapshotFilesResponseBody) *DeleteLiveSnapshotFilesResponse
	GetBody() *DeleteLiveSnapshotFilesResponseBody
}

type DeleteLiveSnapshotFilesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveSnapshotFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveSnapshotFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveSnapshotFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveSnapshotFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveSnapshotFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveSnapshotFilesResponse) GetBody() *DeleteLiveSnapshotFilesResponseBody {
	return s.Body
}

func (s *DeleteLiveSnapshotFilesResponse) SetHeaders(v map[string]*string) *DeleteLiveSnapshotFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveSnapshotFilesResponse) SetStatusCode(v int32) *DeleteLiveSnapshotFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveSnapshotFilesResponse) SetBody(v *DeleteLiveSnapshotFilesResponseBody) *DeleteLiveSnapshotFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveSnapshotFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
