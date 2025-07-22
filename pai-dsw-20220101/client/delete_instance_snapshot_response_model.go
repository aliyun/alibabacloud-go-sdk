// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceSnapshotResponseBody) *DeleteInstanceSnapshotResponse
	GetBody() *DeleteInstanceSnapshotResponseBody
}

type DeleteInstanceSnapshotResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceSnapshotResponse) GetBody() *DeleteInstanceSnapshotResponseBody {
	return s.Body
}

func (s *DeleteInstanceSnapshotResponse) SetHeaders(v map[string]*string) *DeleteInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceSnapshotResponse) SetStatusCode(v int32) *DeleteInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceSnapshotResponse) SetBody(v *DeleteInstanceSnapshotResponseBody) *DeleteInstanceSnapshotResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceSnapshotResponse) Validate() error {
	return dara.Validate(s)
}
