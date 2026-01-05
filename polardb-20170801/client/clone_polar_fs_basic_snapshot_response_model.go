// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClonePolarFsBasicSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClonePolarFsBasicSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClonePolarFsBasicSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *ClonePolarFsBasicSnapshotResponseBody) *ClonePolarFsBasicSnapshotResponse
	GetBody() *ClonePolarFsBasicSnapshotResponseBody
}

type ClonePolarFsBasicSnapshotResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClonePolarFsBasicSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClonePolarFsBasicSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s ClonePolarFsBasicSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ClonePolarFsBasicSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClonePolarFsBasicSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClonePolarFsBasicSnapshotResponse) GetBody() *ClonePolarFsBasicSnapshotResponseBody {
	return s.Body
}

func (s *ClonePolarFsBasicSnapshotResponse) SetHeaders(v map[string]*string) *ClonePolarFsBasicSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponse) SetStatusCode(v int32) *ClonePolarFsBasicSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponse) SetBody(v *ClonePolarFsBasicSnapshotResponseBody) *ClonePolarFsBasicSnapshotResponse {
	s.Body = v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
