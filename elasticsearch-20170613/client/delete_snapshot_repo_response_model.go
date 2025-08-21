// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRepoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnapshotRepoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnapshotRepoResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnapshotRepoResponseBody) *DeleteSnapshotRepoResponse
	GetBody() *DeleteSnapshotRepoResponseBody
}

type DeleteSnapshotRepoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotRepoResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRepoResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRepoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnapshotRepoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnapshotRepoResponse) GetBody() *DeleteSnapshotRepoResponseBody {
	return s.Body
}

func (s *DeleteSnapshotRepoResponse) SetHeaders(v map[string]*string) *DeleteSnapshotRepoResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotRepoResponse) SetStatusCode(v int32) *DeleteSnapshotRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotRepoResponse) SetBody(v *DeleteSnapshotRepoResponseBody) *DeleteSnapshotRepoResponse {
	s.Body = v
	return s
}

func (s *DeleteSnapshotRepoResponse) Validate() error {
	return dara.Validate(s)
}
