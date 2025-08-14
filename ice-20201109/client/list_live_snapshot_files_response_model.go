// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveSnapshotFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveSnapshotFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveSnapshotFilesResponseBody) *ListLiveSnapshotFilesResponse
	GetBody() *ListLiveSnapshotFilesResponseBody
}

type ListLiveSnapshotFilesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveSnapshotFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveSnapshotFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotFilesResponse) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveSnapshotFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveSnapshotFilesResponse) GetBody() *ListLiveSnapshotFilesResponseBody {
	return s.Body
}

func (s *ListLiveSnapshotFilesResponse) SetHeaders(v map[string]*string) *ListLiveSnapshotFilesResponse {
	s.Headers = v
	return s
}

func (s *ListLiveSnapshotFilesResponse) SetStatusCode(v int32) *ListLiveSnapshotFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveSnapshotFilesResponse) SetBody(v *ListLiveSnapshotFilesResponseBody) *ListLiveSnapshotFilesResponse {
	s.Body = v
	return s
}

func (s *ListLiveSnapshotFilesResponse) Validate() error {
	return dara.Validate(s)
}
