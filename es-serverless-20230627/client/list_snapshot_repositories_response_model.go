// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotRepositoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSnapshotRepositoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSnapshotRepositoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListSnapshotRepositoriesResponseBody) *ListSnapshotRepositoriesResponse
	GetBody() *ListSnapshotRepositoriesResponseBody
}

type ListSnapshotRepositoriesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotRepositoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotRepositoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSnapshotRepositoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSnapshotRepositoriesResponse) GetBody() *ListSnapshotRepositoriesResponseBody {
	return s.Body
}

func (s *ListSnapshotRepositoriesResponse) SetHeaders(v map[string]*string) *ListSnapshotRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotRepositoriesResponse) SetStatusCode(v int32) *ListSnapshotRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotRepositoriesResponse) SetBody(v *ListSnapshotRepositoriesResponseBody) *ListSnapshotRepositoriesResponse {
	s.Body = v
	return s
}

func (s *ListSnapshotRepositoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
