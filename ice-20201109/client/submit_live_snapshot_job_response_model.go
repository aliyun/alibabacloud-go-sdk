// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLiveSnapshotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitLiveSnapshotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitLiveSnapshotJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitLiveSnapshotJobResponseBody) *SubmitLiveSnapshotJobResponse
	GetBody() *SubmitLiveSnapshotJobResponseBody
}

type SubmitLiveSnapshotJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitLiveSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitLiveSnapshotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitLiveSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitLiveSnapshotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitLiveSnapshotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitLiveSnapshotJobResponse) GetBody() *SubmitLiveSnapshotJobResponseBody {
	return s.Body
}

func (s *SubmitLiveSnapshotJobResponse) SetHeaders(v map[string]*string) *SubmitLiveSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitLiveSnapshotJobResponse) SetStatusCode(v int32) *SubmitLiveSnapshotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitLiveSnapshotJobResponse) SetBody(v *SubmitLiveSnapshotJobResponseBody) *SubmitLiveSnapshotJobResponse {
	s.Body = v
	return s
}

func (s *SubmitLiveSnapshotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
