// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlternativeSnapshotReposResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAlternativeSnapshotReposResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAlternativeSnapshotReposResponse
	GetStatusCode() *int32
	SetBody(v *ListAlternativeSnapshotReposResponseBody) *ListAlternativeSnapshotReposResponse
	GetBody() *ListAlternativeSnapshotReposResponseBody
}

type ListAlternativeSnapshotReposResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAlternativeSnapshotReposResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAlternativeSnapshotReposResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAlternativeSnapshotReposResponse) GoString() string {
	return s.String()
}

func (s *ListAlternativeSnapshotReposResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAlternativeSnapshotReposResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAlternativeSnapshotReposResponse) GetBody() *ListAlternativeSnapshotReposResponseBody {
	return s.Body
}

func (s *ListAlternativeSnapshotReposResponse) SetHeaders(v map[string]*string) *ListAlternativeSnapshotReposResponse {
	s.Headers = v
	return s
}

func (s *ListAlternativeSnapshotReposResponse) SetStatusCode(v int32) *ListAlternativeSnapshotReposResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlternativeSnapshotReposResponse) SetBody(v *ListAlternativeSnapshotReposResponseBody) *ListAlternativeSnapshotReposResponse {
	s.Body = v
	return s
}

func (s *ListAlternativeSnapshotReposResponse) Validate() error {
	return dara.Validate(s)
}
