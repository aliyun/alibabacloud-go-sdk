// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergSnapshotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIcebergSnapshotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIcebergSnapshotsResponse
	GetStatusCode() *int32
	SetBody(v *ListIcebergSnapshotsResponseBody) *ListIcebergSnapshotsResponse
	GetBody() *ListIcebergSnapshotsResponseBody
}

type ListIcebergSnapshotsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIcebergSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIcebergSnapshotsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListIcebergSnapshotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIcebergSnapshotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIcebergSnapshotsResponse) GetBody() *ListIcebergSnapshotsResponseBody {
	return s.Body
}

func (s *ListIcebergSnapshotsResponse) SetHeaders(v map[string]*string) *ListIcebergSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListIcebergSnapshotsResponse) SetStatusCode(v int32) *ListIcebergSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIcebergSnapshotsResponse) SetBody(v *ListIcebergSnapshotsResponseBody) *ListIcebergSnapshotsResponse {
	s.Body = v
	return s
}

func (s *ListIcebergSnapshotsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
