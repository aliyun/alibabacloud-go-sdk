// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSnapshotsResponseBody) *DeleteSnapshotsResponse
	GetBody() *DeleteSnapshotsResponseBody
}

type DeleteSnapshotsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSnapshotsResponse) GetBody() *DeleteSnapshotsResponseBody {
	return s.Body
}

func (s *DeleteSnapshotsResponse) SetHeaders(v map[string]*string) *DeleteSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotsResponse) SetStatusCode(v int32) *DeleteSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotsResponse) SetBody(v *DeleteSnapshotsResponseBody) *DeleteSnapshotsResponse {
	s.Body = v
	return s
}

func (s *DeleteSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
