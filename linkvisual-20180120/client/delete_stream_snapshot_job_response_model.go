// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStreamSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStreamSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStreamSnapshotJobResponseBody) *DeleteStreamSnapshotJobResponse
	GetBody() *DeleteStreamSnapshotJobResponseBody
}

type DeleteStreamSnapshotJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStreamSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStreamSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStreamSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStreamSnapshotJobResponse) GetBody() *DeleteStreamSnapshotJobResponseBody {
	return s.Body
}

func (s *DeleteStreamSnapshotJobResponse) SetHeaders(v map[string]*string) *DeleteStreamSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamSnapshotJobResponse) SetStatusCode(v int32) *DeleteStreamSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStreamSnapshotJobResponse) SetBody(v *DeleteStreamSnapshotJobResponseBody) *DeleteStreamSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *DeleteStreamSnapshotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
