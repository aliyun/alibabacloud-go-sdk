// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceSnapshotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceSnapshotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceSnapshotResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceSnapshotResponseBody) *ListInstanceSnapshotResponse
	GetBody() *ListInstanceSnapshotResponseBody
}

type ListInstanceSnapshotResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceSnapshotResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceSnapshotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceSnapshotResponse) GetBody() *ListInstanceSnapshotResponseBody {
	return s.Body
}

func (s *ListInstanceSnapshotResponse) SetHeaders(v map[string]*string) *ListInstanceSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceSnapshotResponse) SetStatusCode(v int32) *ListInstanceSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceSnapshotResponse) SetBody(v *ListInstanceSnapshotResponseBody) *ListInstanceSnapshotResponse {
	s.Body = v
	return s
}

func (s *ListInstanceSnapshotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
