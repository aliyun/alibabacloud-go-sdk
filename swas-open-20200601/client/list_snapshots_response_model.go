// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse
	GetBody() *ListSnapshotsResponseBody
}

type ListSnapshotsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSnapshotsResponse) GetBody() *ListSnapshotsResponseBody {
	return s.Body
}

func (s *ListSnapshotsResponse) SetHeaders(v map[string]*string) *ListSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotsResponse) SetStatusCode(v int32) *ListSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotsResponse) SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse {
	s.Body = v
	return s
}

func (s *ListSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
