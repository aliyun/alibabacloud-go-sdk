// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCSnapshotResponseBody) *DeleteRCSnapshotResponse
	GetBody() *DeleteRCSnapshotResponseBody
}

type DeleteRCSnapshotResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCSnapshotResponse) GetBody() *DeleteRCSnapshotResponseBody {
	return s.Body
}

func (s *DeleteRCSnapshotResponse) SetHeaders(v map[string]*string) *DeleteRCSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCSnapshotResponse) SetStatusCode(v int32) *DeleteRCSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCSnapshotResponse) SetBody(v *DeleteRCSnapshotResponseBody) *DeleteRCSnapshotResponse {
	s.Body = v
	return s
}

func (s *DeleteRCSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
