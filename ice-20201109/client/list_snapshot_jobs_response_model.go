// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSnapshotJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSnapshotJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListSnapshotJobsResponseBody) *ListSnapshotJobsResponse
	GetBody() *ListSnapshotJobsResponseBody
}

type ListSnapshotJobsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSnapshotJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSnapshotJobsResponse) GetBody() *ListSnapshotJobsResponseBody {
	return s.Body
}

func (s *ListSnapshotJobsResponse) SetHeaders(v map[string]*string) *ListSnapshotJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotJobsResponse) SetStatusCode(v int32) *ListSnapshotJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotJobsResponse) SetBody(v *ListSnapshotJobsResponseBody) *ListSnapshotJobsResponse {
	s.Body = v
	return s
}

func (s *ListSnapshotJobsResponse) Validate() error {
	return dara.Validate(s)
}
