// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnapshotFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnapshotFilesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnapshotFilesResponseBody) *DeleteSnapshotFilesResponse
	GetBody() *DeleteSnapshotFilesResponseBody
}

type DeleteSnapshotFilesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotFilesResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnapshotFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnapshotFilesResponse) GetBody() *DeleteSnapshotFilesResponseBody {
	return s.Body
}

func (s *DeleteSnapshotFilesResponse) SetHeaders(v map[string]*string) *DeleteSnapshotFilesResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotFilesResponse) SetStatusCode(v int32) *DeleteSnapshotFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotFilesResponse) SetBody(v *DeleteSnapshotFilesResponseBody) *DeleteSnapshotFilesResponse {
	s.Body = v
	return s
}

func (s *DeleteSnapshotFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
