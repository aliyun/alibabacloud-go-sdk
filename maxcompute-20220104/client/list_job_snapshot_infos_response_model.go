// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobSnapshotInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListJobSnapshotInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListJobSnapshotInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListJobSnapshotInfosResponseBody) *ListJobSnapshotInfosResponse
	GetBody() *ListJobSnapshotInfosResponseBody
}

type ListJobSnapshotInfosResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobSnapshotInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobSnapshotInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListJobSnapshotInfosResponse) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListJobSnapshotInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListJobSnapshotInfosResponse) GetBody() *ListJobSnapshotInfosResponseBody {
	return s.Body
}

func (s *ListJobSnapshotInfosResponse) SetHeaders(v map[string]*string) *ListJobSnapshotInfosResponse {
	s.Headers = v
	return s
}

func (s *ListJobSnapshotInfosResponse) SetStatusCode(v int32) *ListJobSnapshotInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobSnapshotInfosResponse) SetBody(v *ListJobSnapshotInfosResponseBody) *ListJobSnapshotInfosResponse {
	s.Body = v
	return s
}

func (s *ListJobSnapshotInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
