// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySnapshotJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySnapshotJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySnapshotJobListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySnapshotJobListResponseBody) *QuerySnapshotJobListResponse
	GetBody() *QuerySnapshotJobListResponseBody
}

type QuerySnapshotJobListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySnapshotJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySnapshotJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySnapshotJobListResponse) GoString() string {
	return s.String()
}

func (s *QuerySnapshotJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySnapshotJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySnapshotJobListResponse) GetBody() *QuerySnapshotJobListResponseBody {
	return s.Body
}

func (s *QuerySnapshotJobListResponse) SetHeaders(v map[string]*string) *QuerySnapshotJobListResponse {
	s.Headers = v
	return s
}

func (s *QuerySnapshotJobListResponse) SetStatusCode(v int32) *QuerySnapshotJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySnapshotJobListResponse) SetBody(v *QuerySnapshotJobListResponseBody) *QuerySnapshotJobListResponse {
	s.Body = v
	return s
}

func (s *QuerySnapshotJobListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
