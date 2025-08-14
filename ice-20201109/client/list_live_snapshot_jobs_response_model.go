// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveSnapshotJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLiveSnapshotJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLiveSnapshotJobsResponse
	GetStatusCode() *int32
	SetBody(v *ListLiveSnapshotJobsResponseBody) *ListLiveSnapshotJobsResponse
	GetBody() *ListLiveSnapshotJobsResponseBody
}

type ListLiveSnapshotJobsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLiveSnapshotJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLiveSnapshotJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLiveSnapshotJobsResponse) GoString() string {
	return s.String()
}

func (s *ListLiveSnapshotJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLiveSnapshotJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLiveSnapshotJobsResponse) GetBody() *ListLiveSnapshotJobsResponseBody {
	return s.Body
}

func (s *ListLiveSnapshotJobsResponse) SetHeaders(v map[string]*string) *ListLiveSnapshotJobsResponse {
	s.Headers = v
	return s
}

func (s *ListLiveSnapshotJobsResponse) SetStatusCode(v int32) *ListLiveSnapshotJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveSnapshotJobsResponse) SetBody(v *ListLiveSnapshotJobsResponseBody) *ListLiveSnapshotJobsResponse {
	s.Body = v
	return s
}

func (s *ListLiveSnapshotJobsResponse) Validate() error {
	return dara.Validate(s)
}
