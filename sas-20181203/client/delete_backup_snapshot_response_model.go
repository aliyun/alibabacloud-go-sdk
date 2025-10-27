// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupSnapshotResponseBody) *DeleteBackupSnapshotResponse
	GetBody() *DeleteBackupSnapshotResponseBody
}

type DeleteBackupSnapshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupSnapshotResponse) GetBody() *DeleteBackupSnapshotResponseBody {
	return s.Body
}

func (s *DeleteBackupSnapshotResponse) SetHeaders(v map[string]*string) *DeleteBackupSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupSnapshotResponse) SetStatusCode(v int32) *DeleteBackupSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupSnapshotResponse) SetBody(v *DeleteBackupSnapshotResponseBody) *DeleteBackupSnapshotResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
