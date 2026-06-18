// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *RestoreSnapshotResponseBody) *RestoreSnapshotResponse
	GetBody() *RestoreSnapshotResponseBody
}

type RestoreSnapshotResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreSnapshotResponse) GoString() string {
	return s.String()
}

func (s *RestoreSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreSnapshotResponse) GetBody() *RestoreSnapshotResponseBody {
	return s.Body
}

func (s *RestoreSnapshotResponse) SetHeaders(v map[string]*string) *RestoreSnapshotResponse {
	s.Headers = v
	return s
}

func (s *RestoreSnapshotResponse) SetStatusCode(v int32) *RestoreSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreSnapshotResponse) SetBody(v *RestoreSnapshotResponseBody) *RestoreSnapshotResponse {
	s.Body = v
	return s
}

func (s *RestoreSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
