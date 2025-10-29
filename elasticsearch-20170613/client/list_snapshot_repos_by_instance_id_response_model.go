// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotReposByInstanceIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSnapshotReposByInstanceIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSnapshotReposByInstanceIdResponse
	GetStatusCode() *int32
	SetBody(v *ListSnapshotReposByInstanceIdResponseBody) *ListSnapshotReposByInstanceIdResponse
	GetBody() *ListSnapshotReposByInstanceIdResponseBody
}

type ListSnapshotReposByInstanceIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotReposByInstanceIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotReposByInstanceIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotReposByInstanceIdResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotReposByInstanceIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSnapshotReposByInstanceIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSnapshotReposByInstanceIdResponse) GetBody() *ListSnapshotReposByInstanceIdResponseBody {
	return s.Body
}

func (s *ListSnapshotReposByInstanceIdResponse) SetHeaders(v map[string]*string) *ListSnapshotReposByInstanceIdResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponse) SetStatusCode(v int32) *ListSnapshotReposByInstanceIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponse) SetBody(v *ListSnapshotReposByInstanceIdResponseBody) *ListSnapshotReposByInstanceIdResponse {
	s.Body = v
	return s
}

func (s *ListSnapshotReposByInstanceIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
